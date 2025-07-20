package adaptor

import (
	"otp/src/pkg/config"
	"otp/src/pkg/log"
	"otp/src/repo"
)

func GetRepoInstance() repo.OTPManagement {
	var otpManagement repo.OTPManagement = nil

	switch config.GetAppConfigInstance().Database {
	case "postgres":
		 CreatePostgresqlDbClient()
		// otpManagement = implementation.GetInstanceOfPostgresOTPManagement(db)

	case "memory":
		// TODO: Implementations

	default:
		log.GetLoggerInstance().Fatal("invalid db type")
	}

	return otpManagement
}
