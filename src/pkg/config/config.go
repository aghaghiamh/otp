package config

type AppConfig struct {
	OTPCacheReservedKey		  string `yaml:"otpCacheReservedKey" envconfig:"OTP_CACHE_RESERVED_KEY"`
	Database                  string `yaml:"database" envconfig:"DB_TYPE"`
	DefaultExpirationInMinute int16  `yaml:"defaultExpirationInMinute" envconfig:"EXPIRATION_IN_MINUTE"`
	AutoMigrationEnable       bool   `yaml:"autoMigrationEnable" envconfig:"AUTO_MIGRATION"`
	SwaggerBaseAddress        string `yaml:"swaggerBaseAddress"`

	Server struct {
		Port int64  `yaml:"port" envconfig:"SERVER_PORT"`
		Host string `yaml:"host" envconfig:"SERVER_HOST"`
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

	Cache struct {
		Host     string `yaml:"host" envconfig:"CACHE_HOST"`
		Port     int    `yaml:"port" envconfig:"CACHE_PORT"`
		Password string `yaml:"password" envconfig:"CACHE_PASSWORD"`
		DB       int  `yaml:"db" envconfig:"CACHE_DB"`
	} `yaml:"cache"`
}
