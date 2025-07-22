package service

import (
	"context"
	"crypto/rand"
	dto "otp/src/controller/httpserver/otpHandler/DTO"
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

func (receiver OTPService) RequestOTP(ctx context.Context, req dto.RequestOTPInput) error {
	// Assumed DDOS attack would be blocked by API Gateway RateLimiter
	if _, err := receiver.otpRepo.Get(context.Background(), *req.MobileNumber); err == nil {
		// TODO: Better handle Logrus fields, perhaps using consts
		log.GetLoggerInstance().WithField("MobileNumber", *req.MobileNumber).Errorf("duplicated OTP Request")

		return &customError.DuplicateError{}
	}

	otpCode := generateRandomCode(6)
	hashedOTP, gErr := bcrypt.GenerateFromPassword([]byte(otpCode), bcrypt.DefaultCost)
	if gErr != nil {
		log.GetLoggerInstance().
			WithField("MobileNumber", *req.MobileNumber).
			WithError(gErr).Errorf("OTP bcrypt code  generation error")
		return gErr
	}

	sErr := receiver.otpRepo.Store(ctx, *req.MobileNumber, string(hashedOTP))
	if sErr != nil {
		log.GetLoggerInstance().
			WithField("MobileNumber", *req.MobileNumber).
			WithError(sErr).Errorf("could not store the generated OTP code into storage")
		return sErr
	}

	log.GetLoggerInstance().WithField("MobileNumber", *req.MobileNumber).Infof("OTP for %s is: %s", *req.MobileNumber, otpCode)
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

func (receiver OTPService) VerifyOTP(ctx context.Context, req dto.VerifyOTPInput) (string, error) {
	codeHash, gErr := receiver.otpRepo.Get(context.Background(), *req.MobileNumber)
	if gErr != nil {
		log.GetLoggerInstance().WithError(gErr).
			WithField("MobileNumber", *req.MobileNumber).
			Error("there is no OTP for this mobile number")

		return "", &customError.OtpNotFoundError{}
	}

	pErr := bcrypt.CompareHashAndPassword([]byte(codeHash), []byte(*req.OtpCode))
	// TODO: Store this malicious request, so then other attempts of user also must be checked
	// to not surpass a limited amount in longer period!
	if pErr != nil {
		log.GetLoggerInstance().WithError(pErr).
			WithField("MobileNumber", *req.MobileNumber).
			Errorf("user provided malicious OTP code: %s", *req.OtpCode)
		return "", &customError.InvalidOtpError{}
	}

	// User Registeration OR Login
	// TODO: Separate the entity.USer and model.User, as it is better to not knowing about the Repo layer.
	// TODO: put the user logic in totally separate service
	user, uErr := receiver.userRepo.GetUserByMobileNumber(ctx, *req.MobileNumber)

	if uErr != nil {
		if uErr == gorm.ErrRecordNotFound {
			user = &model.User{
				MobileNumber: *req.MobileNumber,
			}
			cErr := receiver.userRepo.Register(ctx, user)
			if cErr != nil {
				log.GetLoggerInstance().WithError(uErr).
					WithField("MobileNumber", *req.MobileNumber).
					Error("couldn't register user")
				return "", cErr
			}
		} else {
			return "", uErr
		}
	}

	// TODO: Put all JWT code into the auth service and use it through receiver
	accessToken, aErr := CreateAccessToken(user.ID)
	if aErr != nil {
		// TODO: right now JWT errors better to be treated as server Error, but should be treated otherwise in auth service.
		return "", aErr
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
		// TODO: right now JWT errors better to be treated as server Error, but should be treated otherwise in auth service.
		log.GetLoggerInstance().WithError(signErr).Error("couldn't sign JWT Token")
		return "", signErr
	}

	return tokenString, nil
}
