package main

import (
	_config "example.com/greetings/config"
	"example.com/greetings/rest_route"

	_ "github.com/lib/pq"
)

func main() {
	config := _config.NewConfig()
	route := rest_route.NewRoute()
	route.SetupRoutes(config)
	route.Run("localhost:8080")
}
