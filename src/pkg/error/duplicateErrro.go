package error

import (
	"fmt"
	"net/http"
	"otp/src/pkg/config"
)

type DuplicateError struct{}

func (e *DuplicateError) Error() string {
	return fmt.Sprintf("Repeated OTP request in less than %d minutes is not feasible.",
		config.GetAppConfigInstance().DefaultExpirationInMinute)
}

func (e *DuplicateError) Message() string {
	return fmt.Sprintf("Your previous OTP request was originated in less than %d minutes.",
		config.GetAppConfigInstance().DefaultExpirationInMinute)
}

func (e *DuplicateError) GetErrorCode() int {
	return http.StatusBadRequest
}
