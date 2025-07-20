package adaptor

import (
	"otp/src/pkg/config"
	"otp/src/pkg/log"
	"otp/src/repo"
	"otp/src/repo/implementation"
)

func GetRepoInstance() repo.UserManagement {
	var userManagement repo.UserManagement = nil

	switch config.GetAppConfigInstance().Database {
	case "postgres":
		db := CreatePostgresqlDbClient()
		userManagement = implementation.GetInstanceOfPostgresUserManagement(db)

	case "memory":
		// TODO: Implementations

	default:
		log.GetLoggerInstance().Fatal("invalid db type")
	}

	return userManagement
}
