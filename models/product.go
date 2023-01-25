package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id            uuid.UUID   `json:"id" `
	UserID        uuid.UUID   `json:"user_id"`
	BrandID       uuid.UUID   `json:"brand_id"`
	CategoryID    uuid.UUID   `json:"category_id"`
	Name          string `json:"name"`
	Thumbnail     string `json:"thumbnail"`
	Details       string `json:"details"`
	UnitPrice     string `json:"unit_price"`
	PurchasePrice string `json:"purchase_price"`
	Tax           string `json:"tax"`
	TaxType       string `json:"tax_type"`
	Discount      string `json:"discount"`
	DiscountType  string `json:"discount_type"`
	CurrentStock  string `json:"current_stock"`
	Status        string `json:"status"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (product *Product) Count(db *gorm.DB) int64 {
	var total int64
	
	db.Model(&Product{}).Count(&total)

	return total
}

func (product *Product) Take(db *gorm.DB, limit int, offset int) interface{} {
	var products []Product

	db.Offset(offset).Limit(limit).Find(&products)

	return products
}
