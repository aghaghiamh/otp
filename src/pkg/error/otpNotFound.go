package error

import (
	"fmt"
	"net/http"
	"otp/src/pkg/config"
)

type OtpNotFoundError struct{}

func (e *OtpNotFoundError) Error() string {
	return "No OTP record has been found in the storage"
}

func (e *OtpNotFoundError) Message() string {
	return fmt.Sprintf("No OTP has been requested in previous %d minutes.",
		config.GetAppConfigInstance().DefaultExpirationInMinute)
}

func (e *OtpNotFoundError) GetErrorCode() int {
	return http.StatusNotFound
}
