package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type MedicalHistory struct {
	ID         uuid.UUID     `json:"id"`
	CustomerID uuid.NullUUID `json:"customer_id"`
	PetID      uuid.NullUUID `json:"pet_id"`
	Cost       string        `json:"cost"`
	Date       time.Time     `json:"date"`
	CreatedAt  sql.NullTime  `json:"created_at"`
	UpdatedAt  sql.NullTime  `json:"updated_at"`
}

type MedicalHistoryProduct struct {
	MedicalID uuid.UUID `json:"medical_id"`
	ProductID uuid.UUID `json:"product_id"`
}

type Pet struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Date        time.Time      `json:"date"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type Product struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Cost        string         `json:"cost"`
	Quantity    sql.NullInt32  `json:"quantity"`
	ExpiryDate  sql.NullTime   `json:"expiry_date"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type VeterinaryBill struct {
	ID        int32          `json:"id"`
	Username  sql.NullString `json:"username"`
	Price     sql.NullInt32  `json:"price"`
	CreatedAt sql.NullTime   `json:"created_at"`
}

type VeterinaryMedicine struct {
	ID    int32          `json:"id"`
	Name  sql.NullString `json:"name"`
	Price sql.NullInt32  `json:"price"`
}

type VeterinaryMedicineBill struct {
	BillID     int32 `json:"bill_id"`
	MedicineID int32 `json:"medicine_id"`
}
