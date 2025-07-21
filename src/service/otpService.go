package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"otp/src/model"
	customError "otp/src/pkg/error"
	"otp/src/pkg/log"
	"otp/src/repo"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	jwtSecret = []byte("my-super-secret-key")
	ttl       = 72 * time.Hour
)

type OTPService struct {
	otpRepo  repo.OTPManagement
	userRepo repo.UserManagement
}

func GetInstanceOfOTPService(otpRepo repo.OTPManagement, userRepo repo.UserManagement) *OTPService {
	return &OTPService{
		otpRepo:  otpRepo,
		userRepo: userRepo,
	}
}

func (receiver OTPService) RequestOTP(mobileNumber string) error {
	// Assumed DDOS attack would be blocked by API Gateway RateLimiter
	if _, err := receiver.otpRepo.Get(context.Background(), mobileNumber); err == nil {
		// TODO: Better handle Logrus fields, perhaps using consts
		log.GetLoggerInstance().WithField("MobileNumber", mobileNumber).Errorf("duplicated OTP Request")

		return &customError.DuplicateError{}
	}

	otpCode := generateRandomCode(6)
	hashedOTP, gErr := bcrypt.GenerateFromPassword([]byte(otpCode), bcrypt.DefaultCost)
	if gErr != nil {
		log.GetLoggerInstance().
			WithField("MobileNumber", mobileNumber).
			WithError(gErr).Errorf("OTP bcrypt code  generation error")
		return gErr
	}

	sErr := receiver.otpRepo.Store(context.Background(), mobileNumber, string(hashedOTP))
	if sErr != nil {
		log.GetLoggerInstance().
			WithField("MobileNumber", mobileNumber).
			WithError(sErr).Errorf("could not store the generated OTP code into storage")
		return sErr
	}

	log.GetLoggerInstance().WithField("MobileNumber", mobileNumber).Infof("OTP for %s is: %s", mobileNumber, otpCode)
	return nil
}

// generateRandomCode creates a secure random numeric string of a given length.
func generateRandomCode(length int) string {
	const otpChars = "1234567890"
	buffer := make([]byte, length)
	rand.Read(buffer)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%len(otpChars)]
	}
	return string(buffer)
}

func (receiver OTPService) VerifyOTP(mobileNumber, otpCode string) (string, error) {
	codeHash, gErr := receiver.otpRepo.Get(context.Background(), mobileNumber)
	if gErr != nil {
		log.GetLoggerInstance().Errorf("There is no record with %s mobile number", mobileNumber)
		// TODO: appropriate error Handling
		return "", fmt.Errorf("")
	}

	err := bcrypt.CompareHashAndPassword([]byte(codeHash), []byte(otpCode))
	if err != nil {
		return "", fmt.Errorf("invalid OTP")
	}

	// User Registeration OR Login
	// TODO: Separate the entity.USer and model.User, as it is better to not knowing about the Repo layer.
	// TODO: put the user logic in totally separate service
	user, uErr := receiver.userRepo.GetUserByMobileNumber(context.Background(), mobileNumber)

	if uErr != nil {
		if uErr == gorm.ErrRecordNotFound {
			user = &model.User{
				MobileNumber: mobileNumber,
			}
			cErr := receiver.userRepo.Register(context.Background(), user)
			if cErr != nil {
				return "", cErr
			}
		} else {
			return "", uErr
		}
	}

	// TODO: Put all JWT code into the auth service and use it through receiver
	accessToken, aErr := CreateAccessToken(user.ID)
	if aErr != nil {
		return "", fmt.Errorf("couldn't create JWT Token")
	}

	return accessToken, nil
}

type Claims struct {
	Subject string `json:"subject"`
	UserID  uint   `json:"user_id"`
	jwt.RegisteredClaims
}

func (c *Claims) Valid() {

	return
}

func CreateAccessToken(userID uint) (string, error) {
	const op = "authservice.createToken"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&Claims{
			Subject: "AC",
			UserID:  userID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: &jwt.NumericDate{time.Now().Add(ttl)},
			},
		})

	tokenString, signErr := token.SignedString(jwtSecret)
	if signErr != nil {

		return "", fmt.Errorf("Couldn't sign JWT Token")
	}

	return tokenString, nil
}
