package rest_route

import (
	_config "example.com/greetings/config"
	"github.com/gin-gonic/gin"
)

const (
	getCustomersPath    = "/getCustomers"
	insertCustomersPath = "/insertCustomers"
)

type Route struct {
	gin *gin.Engine
}

func NewRoute() *Route {
	return &Route{
		gin: gin.Default(),
	}
}

func (route *Route) Run(address string) {
	route.gin.Run(address)
}

func (route *Route) SetupRoutes(c *_config.Config) {
	route.gin.GET(getCustomersPath, c.CustomerHandler.GetCustomers)
	route.gin.POST(insertCustomersPath, c.CustomerHandler.InsertCustomers)
}
