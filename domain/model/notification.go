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
	UrlID    int64  `json:"url_id" db:"url_id"`
	Url      string `json:"url" db:"url"`
	IsActive bool   `json:"is_active" db:"is_active"`
}
