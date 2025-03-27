package services

import (
	"database/sql"
	"net/http"

	_const "example.com/greetings/const"
	"example.com/greetings/model"
	_repository "example.com/greetings/repository/customer"
	"github.com/gin-gonic/gin"
)

type ICustomerService interface {
	GetCustomers(c *gin.Context, req model.GetCustomerRequest) ([]model.Customer, int, error)
}

type CustomerService struct {
	repo _repository.ICustomerRepository
}

func NewCustomerService() ICustomerService {
	return &CustomerService{
		repo: _repository.NewCustomerRepository(),
	}
}

func (service *CustomerService) GetCustomers(c *gin.Context, req model.GetCustomerRequest) ([]model.Customer, int, error) {
	db, err := sql.Open(_const.DbDriver, _const.DbSource)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return make([]model.Customer, 0), 0, err
	}

	customers, total, err := service.repo.GetCustomers(db, req)

	return customers, total, err
}
