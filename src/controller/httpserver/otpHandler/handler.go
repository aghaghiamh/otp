package otphandler

import (
	"otp/src/service"
	"otp/src/validator"
)

type Handler struct {
	otpSvc service.OTPService
	validator validator.OTPValidator
}

// TODO: Stick to naming convention of GetInstance..
func New(otpSvc service.OTPService) Handler {
	return Handler{
		otpSvc: otpSvc,
	}
}
