package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

var conf *Configuration

func GetConf() *Configuration {
	if conf != nil {
		return conf
	} else {
		conf, err := LoadConfigurationForEnv(os.Getenv("ENV"))
		if err != nil {
			slog.Error("Unable to load config from env")
			os.Exit(1)
		}

		return conf
	}
}

func LoadConfigurationForEnv(env string) (*Configuration, error) {
	var config *Configuration

	viper.AddConfigPath("./config")
	viper.SetConfigName(fmt.Sprintf("%s.yaml", env))
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Could not read the config file: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshall the config file: %v", err)
	}

	return config, nil
}
