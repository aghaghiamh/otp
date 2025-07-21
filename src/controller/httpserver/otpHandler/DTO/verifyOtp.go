package dto

type UserOTPResponseInfo struct {
	// Name        string `json:"name"`
	MobileNumber string `json:"phone_number"`
}

type AuthTokens struct {
	AccessToken string `json:"access_token"`
	// RefreshToken string `json:"refresh_token"`
}

type VerifyOTPInput struct {
	MobileNumber *string `json:"mobile_number"`
	OtpCode      *string `json:"otp_code"`
}

type VerifyOTPOutput struct {
	UserInfo   UserOTPResponseInfo
	AuthTokens AuthTokens
}
