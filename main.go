package main

import (
	_config "example.com/greetings/config"
	_env_handler "example.com/greetings/env_handler"
	_rest_route "example.com/greetings/rest_route"
	_ "github.com/lib/pq"
)

func main() {
	env := _env_handler.GetEnv()
	config := _config.NewConfig(env)
	route := _rest_route.NewRoute()
	route.SetupRoutes(config)
	route.Run("localhost:8080")
}
