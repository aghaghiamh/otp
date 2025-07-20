package repo

import (
	"context"
	"otp/src/model"
)

type UserManagement interface {
	GetUserByMobileNumber(ctx context.Context, mobileNumber string) (*model.User, error)
	Register(ctx context.Context, user *model.User) error
}
