package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/project/notif-project/domain/model"
	"github.com/project/notif-project/pkg/database"
)

// BrandRepository manages database operations for brand.
type NotifRepository interface {
	GenerateKey(merchantID int64, key string) error
	CheckUrlExistence(form model.Url) (int64, error)
	InsertUrl(form model.Url) error
	GetUrlDetail(urlID int64) (model.GetUrlDetailRes, error)
	UrlToggleStatus(urlID int64) error
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

// GenerateKey Generate a new key for merchant into the database.
func (r *notifRepoImpl) GenerateKey(merchantID int64, key string) error {
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

func (r *notifRepoImpl) CheckUrlExistence(form model.Url) (int64, error) {
	var count int64

	err := r.db.QueryRowx(`
		select count(1)
		from notification.urls u
		where
			u.merchant_id = $1
		and
			(u.url = $2 or u.notification_type = $3)
	`, form.MerchantID, form.Url, form.NotificationType).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *notifRepoImpl) InsertUrl(form model.Url) error {
	_, err := r.db.Exec(`
		INSERT INTO notification.urls (merchant_id, url, notification_type)
		VALUES ($1, $2, $3)
		returning url_id
	`, form.MerchantID, form.Url, form.NotificationType)
	if err != nil {
		return err
	}

	return err
}

func (r *notifRepoImpl) GetUrlDetail(urlID int64) (model.GetUrlDetailRes, error) {
	var res model.GetUrlDetailRes

	err := r.db.QueryRowx(`
		select u.merchant_id, u.url, u.notification_type, u.is_active
		from notification.urls u
		where u.url_id = $1
	`, urlID).StructScan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *notifRepoImpl) UrlToggleStatus(urlID int64) error {
	_, err := r.db.Exec(`
		update notification.urls
		set is_active = not is_active
		where url_id = $1
		`, urlID)
	if err != nil {
		return err
	}

	return err
}
