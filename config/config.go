package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	conf *Configuration
)

func init() {
	once.Do(func() {
		env, ok := os.LookupEnv("ENV")
		if !ok {
			panic("'ENV' environment variable is missing!")
		}

		var config *Configuration

		viper.AddConfigPath("./config")
		viper.SetConfigName(fmt.Sprintf("%s.yaml", env))
		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("Could not read the config file: %v", err))
		}

		if err := viper.Unmarshal(&config); err != nil {
			panic(fmt.Errorf("Could not unmarshal the config file: %v", err))
		}

		conf = config
	})
}

func Conf() *Configuration {
	return conf
}
