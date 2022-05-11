package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/richardsahvic/jamtangan/domain/model"
	"github.com/richardsahvic/jamtangan/pkg/logger"
	"github.com/richardsahvic/jamtangan/service"
)

// TransactionHandler defines dependencies for TransactionHandler.
type TransactionHandler struct {
	transactionService service.TransactionService
}

// NewTransactionhandler returns new instance of TransactionHandler.
func NewTransactionhandler() *TransactionHandler {
	return &TransactionHandler{}
}

// SetTransactionService injects transaction's service for TransactionHandler.
func (h *TransactionHandler) SetTransactionService(service service.TransactionService) *TransactionHandler {
	h.transactionService = service
	return h
}

// Validate validates if all dependency for TransactionHandler is complete.
func (h *TransactionHandler) Validate() *TransactionHandler {
	if h.transactionService == nil {
		log.Panic("Transaction handler need transaction service")
	}
	return h
}

// Transaction handles endpoint with prefix /order
func (h *TransactionHandler) Transaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLoggerContext(ctx, "handler", "Product")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))
	log.Info(fmt.Sprintf("%+v", r))

	w.Header().Set("Content-Type", "application/json")

	var httpCode int
	var resp interface{}

	if r.Method == http.MethodPost {
		var request model.CreateTransactionRequest
		json.Unmarshal(body, &request)

		httpCode, resp = h.transactionService.Create(ctx, request)
	} else if r.Method == http.MethodGet {
		orderID := r.URL.Query().Get("id")

		httpCode, resp = h.transactionService.GetDetail(ctx, orderID)
	} else {
		httpCode = http.StatusMethodNotAllowed
	}

	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(resp)
}
