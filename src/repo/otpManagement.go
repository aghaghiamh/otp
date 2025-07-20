package repo

import (
	"context"
)

type OTPManagement interface {
	Store(ctx context.Context, mobileNumber, otpHash string) error
	Get(ctx context.Context, mobileNumber string) (string, error)
}
