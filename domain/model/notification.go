package model

type Url struct {
	MerchantID       int64  `json:"merchant_id" db:"merchant_id"`
	Url              string `json:"url" db:"url"`
	NotificationType string `json:"notification_type" db:"notification_type"`
}

type GetUrlDetailRes struct {
	MerchantID       int64  `json:"merchant_id" db:"merchant_id"`
	Url              string `json:"url" db:"url"`
	NotificationType string `json:"notification_type" db:"notification_type"`
	IsActive         bool   `json:"is_active" db:"is_active"`
}

type GetMerchantUrlDetail struct {
	UrlID       int64  `json:"url_id" db:"url_id"`
	Url         string `json:"url" db:"url"`
	MerchantKey string `json:"key" db:"key"`
	IsActive    bool   `json:"is_active" db:"is_active"`
}

type CheckOnProsessNotif struct {
	MerchantID    int64 `json:"merchant_id" db:"merchant_id"`
	TransactionID int64 `json:"transaction_id"`
}

type InsertNotifExecution struct {
	MerchantID        int64
	UrlID             int64
	NotificationType  string
	Key               string
	TransactionID     int64
	Amount            float64
	TransactionStatus string
}

type UpdateNotifStatus struct {
	MerchantID         int64
	Key                string
	TransactionID      int64
	NotificationStatus string
}

type SendNotifGoRoutine struct {
	MerchantID        int64
	NotificationType  string
	TransactionID     int64
	Amount            float64
	TransactionStatus string
	CheckSum          string
	IdempotencyKey    string
	Url               string
}
