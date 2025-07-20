package adapter

import "otp/src/pkg/config"

func GetReposInstances() {
	switch config.GetAppConfigInstance().Database {
	case "postgres":
		CreatePostgresqlDbClient()
		// TODO: Implement UserManagement interface and return the interface

	case "memory":
		// TODO: Implementations

	default:
		// TODO: Add Fatal Logger
	}
}
