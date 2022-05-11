package model

// BaseResponse defines the base response of the system.
type BaseResponse struct {
	RawMessage string      `json:"raw_message"`
	ResultData interface{} `json:"data"`
}

type GenerateKeyRequest struct {
	MerchantID int64 `json:"merchant_id"`
}

type NotificationTesterRequest struct {
	UrlID int64 `json:"url_id"`
}

type GenerateKeyResponse struct {
	Key string `json:"key"`
}

type UrlToggleStatusRequest struct {
	UrlID int64 `json:"url_id"`
}
