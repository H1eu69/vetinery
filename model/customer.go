package model

import (
	"database/sql"
	"errors"

	_const "example.com/greetings/const"
	"github.com/google/uuid"
)

type Customer struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	Enabled     bool         `json:"enabled"`
}

type CustomerPet struct {
	CustomerID uuid.UUID `json:"customer_id"`
	PetID      uuid.UUID `json:"pet_id"`
}

type GetCustomerRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Enabled     bool   `form:"enabled"`
	Page        int    `form:"page"`
	PageSize    int    `form:"page_size"`
}

type InsertCustomerRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Enabled     bool   `form:"enabled"`
}

func (req *GetCustomerRequest) Validate() error {
	if req.Page < 1 {
		req.Page = _const.DefaultPage
	}
	if req.PageSize < 1 {
		req.PageSize = _const.DefaultPageSize
	}
	if req.Name == "" {
		return errors.New(_const.ErrBadParams)
	}

	return nil
}
