package config

import (
	_env_handler "example.com/greetings/env_handler"
	"example.com/greetings/handler"
)

type Config struct {
	CustomerHandler handler.ICustomerHandler
	EnvHandler      _env_handler.EnvironmentHandler
}

func NewConfig(env *_env_handler.EnvironmentHandler) *Config {
	return &Config{
		CustomerHandler: handler.NewCustomerHandler(),
		EnvHandler:      *env,
	}
}
