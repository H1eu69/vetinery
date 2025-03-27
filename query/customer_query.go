package query

import (
	"database/sql"
	"example.com/greetings/model"
	"log"
)

func GetAllCustomer(db *sql.DB) []model.Customer {
	query := `SELECT * FROM CUSTOMER`
	row, err := db.Query(query)
	if err != nil {
		log.Fatal("cannot execute query:", err)
	}
	var userrslt []model.Customer
	for row.Next() {
		customer := model.Customer{}
		err := row.Scan(&customer.ID, &customer.Name, &customer.Description, &customer.CreatedAt, &customer.UpdatedAt, &customer.Enabled)
		if err != nil {
			log.Fatal("cannot execute query:", err)
		}
		userrslt = append(userrslt, customer)
	}
	return userrslt

}

func InsertCustomerToDB(db *sql.DB, customer model.Customer) string {
	query := `INSERT INTO Customer ( name, description)
	VALUES ($1, $2) RETURNING ID`
	var id string
	err := db.QueryRow(query, customer.Name, customer.Description).Scan(&id)
	if err != nil {
		log.Fatal("cannot execute query:", err)
	}
	return id
}
