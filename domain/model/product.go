package model

import (
	"database/sql"
	"time"
)

// Product contains details of product.
type Product struct {
	ID        int64        `json:"id" db:"id"`
	SKU       string       `json:"sku" db:"sku"`
	BrandID   int64        `json:"brand_id" db:"brand_id"`
	Stock     int64        `json:"stock" db:"stock"`
	Price     float64      `json:"pric" db:"price"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
