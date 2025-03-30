package repository

import (
	"database/sql"
	"log"
	"strings"

	"example.com/greetings/model"
)

type ICustomerRepository interface {
	GetCustomers(db *sql.DB, req model.GetCustomerRequest) ([]model.Customer, int, error)
	InsertCustomers(db *sql.DB, req model.InsertCustomerRequest) (*model.Customer, error)
}

type CustomerRepository struct {
}

func NewCustomerRepository() ICustomerRepository {
	return &CustomerRepository{}
}

func (repo *CustomerRepository) GetCustomers(db *sql.DB, req model.GetCustomerRequest) ([]model.Customer, int, error) {
	baseQuery := `SELECT * FROM CUSTOMER`
	conditions := []string{}

	if req.Name != "" {
		conditions = append(conditions, "name = '"+req.Name+"'")
	}
	if req.Description != "" {
		conditions = append(conditions, "description = '"+req.Description+"'")
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	row, err := db.Query(baseQuery)
	if err != nil {
		log.Print("cannot execute query:", err)
	}
	customers := make([]model.Customer, 0)

	for row.Next() {
		customer := model.Customer{}
		err := row.Scan(&customer.ID, &customer.Name, &customer.Description, &customer.CreatedAt, &customer.UpdatedAt, &customer.Enabled)
		if err != nil {
			log.Print("cannot execute query:", err)
		}
		customers = append(customers, customer)
	}
	return customers, len(customers), nil
}

func (repo *CustomerRepository) InsertCustomers(db *sql.DB, req model.InsertCustomerRequest) (*model.Customer, error) {
	baseQuery := `INSERT INTO CUSTOMER (name, description, enabled) VALUES ($1, $2, $3)
	RETURNING id, name, description, enabled, created_at, updated_at`
	row := db.QueryRow(baseQuery, req.Name, req.Description, req.Enabled)

	customer := model.Customer{}

	err := row.Scan(&customer.ID, &customer.Name, &customer.Description, &customer.Enabled, &customer.CreatedAt, &customer.UpdatedAt)

	return &customer, err

}
