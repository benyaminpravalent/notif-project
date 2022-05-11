package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/project/notif-project/pkg/database"
)

// BrandRepository manages database operations for brand.
type NotifRepository interface {
	Create(merchantID int64, key string) error
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

// Create creates a new brand into the database.
func (r *notifRepoImpl) Create(merchantID int64, key string) error {
	_, err := r.db.Exec(`
		update notification.merchants
		set key = $1
		where merchant_id = $2
		`, key, merchantID)
	if err != nil {
		return err
	}

	return err
}
