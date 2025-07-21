package dto

type RequestOTPInput struct {
	MobileNumber *string `json:"mobile_number"`
}

type RequestOTPOutput struct {
	Message string `json:"message"`
}
