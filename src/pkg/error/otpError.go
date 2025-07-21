package error

type OTPError interface {
	error
	Message() string
	GetErrorCode() int
}
