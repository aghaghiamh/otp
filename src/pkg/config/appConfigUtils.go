package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

func GetRunningEnv() string {
	env := os.Getenv("environment")
	if env == "" {
		return "local"
	}
	return env
}

func readAppConfig(configAddress *string) *AppConfig {
	var cfg AppConfig
	env := GetRunningEnv()
	readYamlFile(&cfg, configAddress, nil)
	readYamlFile(&cfg, configAddress, &env)
	readEnv(&cfg)
	return &cfg
}

func readYamlFile(cfg *AppConfig, configAddress *string, env *string) {
	baseAddress := "./config/"
	if configAddress != nil {
		baseAddress = *configAddress
	}
	fileAddress := baseAddress + "config.yaml"

	if env != nil && *env != "" {
		fileAddress = baseAddress + "config." + *env + ".yaml"
	}
	f, err := os.Open(fileAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func readEnv(cfg *AppConfig) {
	// Process environment variables into the config struct
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err)
	}
}
