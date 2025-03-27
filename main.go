package main

import (
	"example.com/greetings/https"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	route := gin.Default()
	route.GET("/getCustomers", https.GetCustomers)
	// route.POST("/insertCustomer", https.InsertCustomer)
	route.Run("localhost:8080")
}
