package validator

import (
	"regexp"

	dto "otp/src/controller/httpserver/otpHandler/DTO"
	customError "otp/src/pkg/error"
	"otp/src/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v OTPValidator) ValidateRequestOTP(req dto.RequestOTPInput) (map[string]string, error) {

	err := validation.ValidateStruct(&req,
		validation.Field(
			&req.MobileNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(`^(\(?\+98\)?)?[-\s]?(09)(\d{9})$`)).Error("Mobile Number does not satisfy the valid pattern of `(+98) 09xxxxxxxxx`.")),
	)

	if err != nil {
		fieldErrors := map[string]string{}
		if vErr, ok := err.(validation.Errors); ok {
			for key, val := range vErr {
				if val != nil {
					fieldErrors[key] = val.Error()
				}
				log.GetLoggerInstance().WithError(val)
			}
		}
		return fieldErrors, &customError.InvalidInputError{}
	}

	return nil, nil
}
