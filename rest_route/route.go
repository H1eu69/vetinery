package rest_route

import (
	_config "example.com/greetings/config"
	"github.com/gin-gonic/gin"
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
	route.gin.GET("/getCustomers", c.CustomerHandler.GetCustomers)

}
