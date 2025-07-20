package adapter

import (
	"otp/src/pkg/config"
	"otp/src/pkg/log"
	"otp/src/repo"
	"otp/src/repo/implementation"
)

func GetReposInstances() repo.OTPManagement {
	var otpManagement repo.OTPManagement = nil

	switch config.GetAppConfigInstance().Database {
	case "postgres":
		db := CreatePostgresqlDbClient()
		otpManagement = implementation.GetInstanceOfPostgresOTPManagement(db)

	case "memory":
		// TODO: Implementations

	default:
		log.GetLoggerInstance().Fatal("invalid db type")
	}

	return otpManagement
}
