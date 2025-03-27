package main

import (
	"example.com/greetings/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	config := NewConfig()
	route := gin.Default()
	route.GET("/getCustomers", config.customerHandler.GetCustomers)
	// route.POST("/insertCustomer", https.InsertCustomer)
	route.Run("localhost:8080")
}

type Config struct {
	customerHandler handler.ICustomerHandler
}

func NewConfig() *Config {
	return &Config{
		customerHandler: handler.NewCustomerHandler(),
	}
}
