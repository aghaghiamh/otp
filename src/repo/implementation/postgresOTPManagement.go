package implementation

import (
	"errors"
	"otp/src/model"
	"time"

	"otp/src/pkg/config"
	"otp/src/pkg/log"

	"gorm.io/gorm"
)

type PostgresOTPManagement struct {
	DB *gorm.DB
}

func GetInstanceOfPostgresOTPManagement(db *gorm.DB) *PostgresOTPManagement {
	return &PostgresOTPManagement{
		DB: db,
	}
}

func (pgo *PostgresOTPManagement) Store(otp *model.OTP) error {
	otp.CreatedAt = time.Now()
	otp.ExpiresAt = time.Now().Add(
		time.Duration(config.GetAppConfigInstance().DefaultExpirationInMinute) * time.Minute)
	
	cErr := pgo.DB.Create(otp).Error
	if cErr != nil {
		log.GetLoggerInstance().WithError(cErr).Errorf("failed to insert otp for %s mobile number", otp.MobileNumber)
		return errors.New("couldn't create error!")
	}
	return nil
}
