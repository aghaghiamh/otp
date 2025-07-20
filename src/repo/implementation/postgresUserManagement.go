package implementation

import (
	"context"
	"errors"
	"otp/src/model"

	"gorm.io/gorm"
)

type PostgresUserManagement struct {
	DB *gorm.DB
}

func GetInstanceOfPostgresUserManagement(db *gorm.DB) *PostgresUserManagement {
	return &PostgresUserManagement{
		DB: db,
	}
} 

func (pgu *PostgresUserManagement) GetUserByMobileNumber(ctx context.Context, mobileNumber string) (*model.User, error) {
	var user model.User

    result := pgu.DB.WithContext(ctx).Where("mobile_number = ?", mobileNumber).First(&user)

    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// TODO: Send more general Error, which is not restricted to GORM, as the service layer shouldn't know about implementations of the repo
            return nil, gorm.ErrRecordNotFound
        }
        return nil, result.Error
    }

    return &user, nil
}

func (pgu *PostgresUserManagement) Register(ctx context.Context, user *model.User) error {
    result := pgu.DB.WithContext(ctx).Create(user)
	
    return result.Error
}
