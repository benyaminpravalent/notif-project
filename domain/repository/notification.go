package repository

import (
	"errors"

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
	GetMerchantUrlDetail(merchantID int64, notificationType string) (model.GetMerchantUrlDetail, error)
	CheckOnProsessNotif(form model.CheckOnProsessNotif) (int64, error)
	InsertNotifExecution(form model.InsertNotifExecution) error
	UpdateNotifStatus(form model.UpdateNotifStatus) error
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

func (r *notifRepoImpl) GetMerchantUrlDetail(merchantID int64, notificationType string) (model.GetMerchantUrlDetail, error) {
	var res model.GetMerchantUrlDetail

	err := r.db.QueryRowx(`
		select u.url_id, u.url, u.is_active, m.key
		from notification.urls u
		join notification.merchants m on m.merchant_id = u.merchant_id
		where u.merchant_id = $1 and u.notification_type = $2
		limit 1
	`, merchantID, notificationType).StructScan(&res)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return res, errors.New("Merchant does not have the notification webhook yet")
		}
		return res, err
	}

	return res, nil
}

func (r *notifRepoImpl) CheckOnProsessNotif(form model.CheckOnProsessNotif) (int64, error) {
	var count int64

	err := r.db.QueryRowx(`
		select count(1)
		from notification.notifications n
		where
			n.merchant_id = $1 and n.transaction_id = $2
		and
			n.notification_status = 'failed' or n.notification_status = 'pending'
	`, form.MerchantID, form.TransactionID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *notifRepoImpl) InsertNotifExecution(form model.InsertNotifExecution) error {
	_, err := r.db.Exec(`
		INSERT INTO notification.notifications (
			merchant_id,
			url_id,
			notification_type,
			transaction_id,
			amount,
			transaction_status,
			idempotency_key,
			notification_status
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)
	`,
		form.MerchantID,
		form.UrlID,
		form.NotificationType,
		form.TransactionID,
		form.Amount,
		form.TransactionStatus,
		form.Key,
		"pending",
	)
	if err != nil {
		return err
	}

	return err
}

func (r *notifRepoImpl) UpdateNotifStatus(form model.UpdateNotifStatus) error {
	_, err := r.db.Exec(`
		update notification.notifications
		set
			notification_status = $1
		where
			merchant_id = $2 and idempotency_key = $3 and transaction_id = $4
		`, form.NotificationStatus, form.MerchantID, form.Key, form.TransactionID)
	if err != nil {
		return err
	}

	return err
}
