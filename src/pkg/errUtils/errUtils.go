package errutils

import (
	customError "otp/src/pkg/error"
)

func GetStatusCode(err error) int {
	errorCode := 500
	switch e := err.(type) {
	case customError.OTPError:
		errorCode = e.GetErrorCode()
	}
	return errorCode
}

func GenerateErrorMessage(err error) string {
	defaultError := "There is a problem in processing your information."
	switch e := err.(type) {
	case customError.OTPError:
		return e.Message()
	default:
		return defaultError
	}
}