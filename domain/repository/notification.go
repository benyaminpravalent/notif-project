package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/project/notif-project/pkg/database"
)

// BrandRepository manages database operations for brand.
type NotifRepository interface {
}

type notifRepoImpl struct {
	db *sqlx.DB
}

// NewBrandRepository returns new instance of brandRepoImpl.
func NewNotifRepository() *notifRepoImpl {
	return &notifRepoImpl{
		db: database.DB,
	}
}
