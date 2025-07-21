package dto

type VerifyOTPInputDTO struct {
	MobileNumber *string `json:"mobile_number"`
	OtpCode      *string `json:"otp_code"`
}
