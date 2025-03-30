package config

import (
	"example.com/greetings/handler"
)

type Config struct {
	CustomerHandler handler.ICustomerHandler
}

func NewConfig() *Config {
	return &Config{
		CustomerHandler: handler.NewCustomerHandler(),
	}
}
