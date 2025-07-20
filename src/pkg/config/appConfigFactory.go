package config

import "sync"

var (
	createOnce sync.Once
	config     *AppConfig
)

func GetAppConfigInstance(configAddress ...string) *AppConfig {
	var configDirectory *string = nil
	if len(configAddress) != 0 {
		configDirectory = &configAddress[0]
	}
	createOnce.Do(func() {
		config = readAppConfig(configDirectory)
	})
	return config
}