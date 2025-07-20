package repo

import "otp/src/model"

type OTPManagement interface {
	Store(otp *model.OTP) error
}
