package service

import (
	"crypto/rand"
	"otp/src/model"
	"otp/src/pkg/log"
	"otp/src/repo"

	"golang.org/x/crypto/bcrypt"
)

type OTPService struct {
	otpRepo repo.OTPManagement
}

func GetInstanceOfOTPService(otpRepo repo.OTPManagement) *OTPService {
	return &OTPService{otpRepo: otpRepo}
}

func (receiver OTPService) RequestOTP(mobileNumber string) error {
	// TODO: rate limit user

	otpCode := generateRandomCode(6)

	hashedOTP, err := bcrypt.GenerateFromPassword([]byte(otpCode), bcrypt.DefaultCost)
	if err != nil {
		// TODO: handle properly
		return err
	}

	sErr := receiver.otpRepo.Store(&model.OTP{
		MobileNumber: mobileNumber,
		CodeHash:     string(hashedOTP),
	})
	if sErr != nil {
		// TODO: Handle Error
	}

	log.GetLoggerInstance().Infof("OTP for %s is: %s", mobileNumber, otpCode)

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
