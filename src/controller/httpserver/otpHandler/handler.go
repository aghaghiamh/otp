package otphandler

import (
	"otp/src/service"
)

type Handler struct {
	otpSvc service.OTPService
	// TODO: ADD Validation
}

// TODO: Stick to naming convention of GetInstance..
func New(otpSvc service.OTPService) Handler {
	return Handler{
		otpSvc: otpSvc,
	}
}
