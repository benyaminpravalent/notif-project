package model

import (
	"time"
)

// Brand contains details of brand.
type Brand struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	// UpdatedAt sql.NullTime `db:"updated_at"`
	// DeletedAt sql.NullTime `db:"deleted_at"`
}
