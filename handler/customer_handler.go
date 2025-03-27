package handler

import (
	"log"
	"net/http"

	response "example.com/greetings/https/response"
	"example.com/greetings/model"
	services "example.com/greetings/services/customer"
	"github.com/gin-gonic/gin"
)

type ICustomerHandler interface {
	GetCustomers(c *gin.Context)
}

type CustomerHandler struct {
	Service services.ICustomerService
}

func NewCustomerHandler() ICustomerHandler {
	return &CustomerHandler{Service: services.NewCustomerService()}
}

func (h *CustomerHandler) GetCustomers(c *gin.Context) {
	var getCustomerRequest model.GetCustomerRequest

	if err := c.ShouldBindQuery(&getCustomerRequest); err != nil {
		log.Print("cannot read body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if reqErr := getCustomerRequest.Validate(); reqErr != nil {
		log.Print("bad params: ", reqErr)
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
	}

	customers, total, err := services.NewCustomerService().GetCustomers(c, getCustomerRequest)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
		log.Print("cannot get customers:", err)
		return
	}
	response.SuccessResponse(c, "success", map[string]any{"list": customers, "total": total})
}

// func InsertCustomer(c *gin.Context) {
// 	var customer model.Customer

// 	if err := c.BindJSON(&customer); err != nil {
// 		log.Print("cannot read body:", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db, err := sql.Open(dbDriver, dbSource)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		log.Print("cannot connect to db:", err)
// 	}

// 	var id = query.InsertCustomerToDB(db, customer)

// 	c.JSON(http.StatusOK, gin.H{"id": id})
// }

// func UpdateCustomer(c *gin.Context) {
// 	var getCustomer model.GetCustomerRequest

// 	if err := c.BindJSON(&getCustomer); err != nil {
// 		log.Print("cannot read body:", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	reqErr := getCustomer.Validate()
// 	if reqErr != nil {
// 		log.Print("bad params", reqErr)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
// 		return
// 	}
// 	db, err := sql.Open(dbDriver, dbSource)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		log.Print("cannot connect to db:", err)
// 	}

// 	var id = query.InsertCustomerToDB(db, getCustomer)

// 	c.JSON(http.StatusOK, gin.H{"id": id})
// }
