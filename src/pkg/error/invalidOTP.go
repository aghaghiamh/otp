package error

import (
	"net/http"
)

type InvalidOtpError struct{}

func (e *InvalidOtpError) Error() string {
	return "User provided OTP mismatches the server generated one."
}

func (e *InvalidOtpError) Message() string {
	return "The OTP code you provided, doesn't match that one the server has been generated!"
}

func (e *InvalidOtpError) GetErrorCode() int {
	return http.StatusBadRequest
}
