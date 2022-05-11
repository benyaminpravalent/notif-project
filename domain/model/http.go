package model

// CreateBrandRequest defines request to create brand.
type CreateBrandRequest struct {
	Name string `json:"name"`
}

// CreateProductRequest defines request to create product.
type CreateProductRequest struct {
	BrandID int64   `json:"brand_id"`
	SKU     string  `json:"sku"`
	Stock   int64   `json:"stock"`
	Price   float64 `json:"price"`
}

// BaseResponse defines the base response of the system.
type BaseResponse struct {
	RawMessage string      `json:"raw_message"`
	ResultData interface{} `json:"data"`
}

// CreateBrandResponse defines response to create brand.
type CreateBrandResponse struct {
	ID int64 `json:"id"`
}

// CreateProductResponse defines response to create product.
type CreateProductResponse struct {
	ID int64 `json:"id"`
}

// GetProductResponse defines response to get product.
type GetProductResponse struct {
	ID      int64   `json:"id"`
	BrandID int64   `json:"brand_id"`
	SKU     string  `json:"sku"`
	Stock   int64   `json:"stock"`
	Price   float64 `json:"price"`
}

// GetProductByBrandIDResponse defines response to get product by brand.
type GetProductByBrandIDResponse struct {
	Products []*Product `json:"products"`
}

// TransactionItem defines the items in transactions.
type TransactionItem struct {
	SKU      string  `json:"sku"`
	Quantity int64   `json:"quantity"`
	Subtotal float64 `json:"subtotal"`
}

// CreateTransactionRequest defines request to create transaction.
type CreateTransactionRequest struct {
	Items []TransactionItem `json:"items"`
}

// CreateTransactionResponse defines response to create transaction.
type CreateTransactionResponse struct {
	OrderID    string  `json:"order_id"`
	TotalPrice float64 `json:"total_price"`
}

// GetTranscationDetailResponse defines response to get transaction detail.
type GetTranscationDetailResponse struct {
	OrderID     string            `json:"order_id"`
	Items       []TransactionItem `json:"items"`
	TotalAmount float64           `json:"total_amount"`
}

type GenerateKeyRequest struct {
	MerchantID int64 `json:"merchant_id"`
}

type GenerateKeyResponse struct {
	Key string `json:"key"`
}
