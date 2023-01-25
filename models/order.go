package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	Id                uuid.UUID `json:"id"`
	UserId            uuid.UUID `json:"user_id"`
	DeliveryAddressId string    `json:"delivery_address_id"`
	OrderAmount       string    `json:"order_amount" `
	OrderStatus       string    `json:"order_status"`
	PaymentStatus     string    `json:"payment_status"`
	TotalTaxAmount    string    `json:"total_tax_amount"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	OrderItems        []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id          uuid.UUID `json:"id"`
	ProductId   uuid.UUID `json:"product_id"`
	OrderId     uuid.UUID `json:"order_id"`
	Price       float64   `json:"price"`
	ProductName string    `json:"product_name"`
	Image       string    `json:"image"`
	Discount    string    `json:"discount"`
	Quantity    float64   `json:"quantity"`
	TaxAmount   string    `json:"tax_amount"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
