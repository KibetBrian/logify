package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type Address struct {
	Id          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	ContactName string    `json:"contact_name"`
	AddressType string    `json:"address_type"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	Zip         string    `json:"zip"`
	Phone       string    `json:"phone"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
