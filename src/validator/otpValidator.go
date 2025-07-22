package validator

const (
	PhoneNumberRegex = `^(\(?\+98\)?)?[-\s]?(09)(\d{9})$`
)

type OTPRepo interface {
}

type OTPValidator struct {
}

func New(repo OTPRepo) OTPValidator {
	return OTPValidator{}
}
