package model

import (
	"database/sql"
	"time"
)

// Transaction contains detail of transaction.
type Transaction struct {
	ID        int64        `json:"id" db:"id"`
	SKU       string       `json:"sku" db:"sku"`
	Quantity  int64        `json:"quantity" db:"quantity"`
	OrderID   string       `json:"order_id" db:"order_id"`
	Subtotal  float64      `json:"subtotal" db:"subtotal"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
