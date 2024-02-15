package configs

import "mono.thienhang.com/pkg/config"

type (
	// main app configuration
	AppConfig struct {
		config.BaseConfig
		Http string `mapstructure:"http"`
	}
)
