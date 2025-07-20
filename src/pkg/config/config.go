package config

type AppConfig struct {
	Database               string `yaml:"database" envconfig:"DB_TYPE"`
	DefaultExpirationInMinute int16  `yaml:"defaultExpirationInDay" envconfig:"EXPIRATION_IN_MINUTE"`
	AutoMigrationEnable    bool   `yaml:"autoMigrationEnable" envconfig:"AUTO_MIGRATION"`
	SwaggerBaseAddress     string `yaml:"swaggerBaseAddress"`

	Server struct {
		Port int64 `yaml:"port" envconfig:"SERVER_PORT"`
	} `yaml:"server"`

	Log struct {
		Level string `yaml:"level" envconfig:"LOG_LEVEL"`
	} `yaml:"log"`

	Postgres struct {
		Username string `yaml:"user" envconfig:"PG_USERNAME"`
		Password string `yaml:"pass" envconfig:"PG_PASSWORD"`
		DB       string `yaml:"db"   envconfig:"PG_DB"`
		Host     string `yaml:"host" envconfig:"PG_HOST"`
		Port     int64  `yaml:"port" envconfig:"PG_PORT"`
	} `yaml:"postgres"`
}
