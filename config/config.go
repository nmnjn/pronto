package config

import "strings"

const (
	PRODUCTION_ENVIRONMENT  = "production"
	STAGING_ENVIRONMENT     = "staging"
	DEVELOPMENT_ENVIRONMENT = "development"
)

type RouterConfig struct {
	Port int16
}

type LoggerConfig struct {
	Level        string
	EnableColor  bool
	EnableCaller bool
}

type ProntoConfig struct {
	Name        string
	Environment string
	RouterConfig
	LoggerConfig
}

func (c *ProntoConfig) Production() bool {
	return strings.ToLower(c.Environment) == PRODUCTION_ENVIRONMENT
}
