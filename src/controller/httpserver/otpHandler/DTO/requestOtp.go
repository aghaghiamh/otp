package dto

type RequestOTPInputDTO struct {
	MobileNumber *string `json:"mobile_number"`
}

type RequestOTPSuccessOutputDTO struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
