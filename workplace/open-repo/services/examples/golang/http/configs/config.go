package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var (
	cfg *AppConfig
)

func InitializeConfig(path string) *AppConfig {
	if path != "" {
		viper.AddConfigPath(path)
	}
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.SetEnvPrefix("app")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	cfg = new(AppConfig)
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err)
	}
	return cfg
}

func (c *AppConfig) LoadAdditional(additional interface{}) {
	if err := viper.Unmarshal(additional); err != nil {
		panic(err)
	}
}

func Get() AppConfig {
	if cfg == nil {
		panic("Config was not initialized")
	}
	return *cfg
}

func overrideWithEnvVars() {
	for _, key := range viper.AllKeys() {
		viper.Set(key, viper.Get(key))
	}
}

func AllSettings() map[string]interface{} {
	if cfg == nil {
		panic("Config is not initialized")
	}
	return viper.AllSettings()
}
