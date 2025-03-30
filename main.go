package main

import (
	_config "example.com/greetings/config"
	_rest_route "example.com/greetings/rest_route"

	_ "github.com/lib/pq"
)

func main() {
	config := _config.NewConfig()
	route := _rest_route.NewRoute()
	route.SetupRoutes(config)
	route.Run("localhost:8080")
}
