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

// BrandHandler defines dependencies for brand handler.
type BrandHandler struct {
	brandService service.BrandService
}

// NewBrandHandler returns new instance of BrandHandler
func NewBrandHandler() *BrandHandler {
	return &BrandHandler{}
}

// SetBrandService injects brand's service for Brandhandler
func (h *BrandHandler) SetBrandService(service service.BrandService) *BrandHandler {
	h.brandService = service
	return h
}

// Validate validates if all dependency for BrandHandler is complete.
func (h *BrandHandler) Validate() *BrandHandler {
	if h.brandService == nil {
		log.Panic("Brand handler need brand service")
	}
	return h
}

// Brand handles endpoint with prefix /brand.
func (h *BrandHandler) Brand(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLoggerContext(ctx, "handler", "Brand")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))
	log.Info(fmt.Sprintf("%+v", r))

	w.Header().Set("Content-Type", "application/json")

	var httpCode int
	var resp interface{}

	if r.Method == http.MethodPost {
		var request model.CreateBrandRequest
		json.Unmarshal(body, &request)

		httpCode, resp = h.brandService.Create(ctx, request)
	} else {
		httpCode = http.StatusMethodNotAllowed
	}

	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(resp)
}
