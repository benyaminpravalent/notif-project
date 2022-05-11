package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/project/notif-project/domain/model"
	"github.com/project/notif-project/pkg/logger"
	"github.com/project/notif-project/service"
)

// ProductHandler defines dependencies for product handler.
type ProductHandler struct {
	productService service.ProductService
}

// NewProductHandler returns new instance of ProductHandler.
func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// SetProductService injects product's service for ProductHandler.
func (h *ProductHandler) SetProductService(service service.ProductService) *ProductHandler {
	h.productService = service
	return h
}

// Validate validates if all dependency for ProductHandler is complete.
func (h *ProductHandler) Validate() *ProductHandler {
	if h.productService == nil {
		log.Panic("Product handler need product service")
	}
	return h
}

// Product handles endpoint with prefix /product
func (h *ProductHandler) Product(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLoggerContext(ctx, "handler", "Product")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))
	log.Info(fmt.Sprintf("%+v", r))

	w.Header().Set("Content-Type", "application/json")

	var httpCode int
	var resp interface{}

	if r.Method == http.MethodPost {
		var request model.CreateProductRequest
		json.Unmarshal(body, &request)

		httpCode, resp = h.productService.Create(ctx, request)
	} else if r.Method == http.MethodGet {
		productID := r.URL.Query().Get("id")

		httpCode, resp = h.productService.GetByID(ctx, productID)
	} else {
		httpCode = http.StatusMethodNotAllowed
	}

	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(resp)
}

// ProductByBrand handles endpoint with prefix /product/brand
func (h *ProductHandler) ProductByBrand(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLoggerContext(ctx, "handler", "ProductByBrand")

	log.Info(fmt.Sprintf("%+v", r))

	w.Header().Set("Content-Type", "application/json")

	var httpCode int
	var resp interface{}

	if r.Method == http.MethodGet {
		brandID := r.URL.Query().Get("id")

		httpCode, resp = h.productService.GetByBrandID(ctx, brandID)
	} else {
		httpCode = http.StatusMethodNotAllowed
	}

	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(resp)
}
