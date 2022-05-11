package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/richardsahvic/jamtangan/domain/model"
	"github.com/richardsahvic/jamtangan/domain/repository"
	"github.com/richardsahvic/jamtangan/pkg/logger"
	"github.com/richardsahvic/jamtangan/pkg/utils"
)

// TransactionService manage logical syntax for transaction.
type TransactionService interface {
	Create(ctx context.Context, request model.CreateTransactionRequest) (int, *model.BaseResponse)
	GetDetail(ctx context.Context, orderID string) (int, *model.BaseResponse)
}

type transactionServiceImpl struct {
	transactionRepo repository.TransactionRepository
}

// NewTransactionService returns new instance of transactionServiceImpl.
func NewTransactionService() *transactionServiceImpl {
	return &transactionServiceImpl{}
}

// SetTransactionRepo injects transaction's repo for transactionServiceImpl
func (s *transactionServiceImpl) SetTransactionRepo(repo repository.TransactionRepository) *transactionServiceImpl {
	s.transactionRepo = repo
	return s
}

// Validate validates if all dependency for transactionServiceImpl is complete.
func (s *transactionServiceImpl) Validate() *transactionServiceImpl {
	if s.transactionRepo == nil {
		log.Panic("Transaction service need transaction repository")
	}
	return s
}

// Create creates a new transaction and store it into the database.
func (s *transactionServiceImpl) Create(ctx context.Context, request model.CreateTransactionRequest) (int, *model.BaseResponse) {
	// validate request
	if len(request.Items) == 0 {
		return utils.RequestRequired("items")
	}

	log := logger.GetLoggerContext(ctx, "service", "Create")

	orderID := utils.GenerateOrderID()

	var totalPrice float64

	order := make([]model.Transaction, 0)
	for _, item := range request.Items {
		order = append(order, model.Transaction{
			OrderID:  orderID,
			SKU:      item.SKU,
			Quantity: item.Quantity,
			Subtotal: item.Subtotal,
		})

		totalPrice += item.Subtotal
	}

	err := s.transactionRepo.InsertList(order)
	if err != nil {
		log.Error(fmt.Sprintf("failed to create transaction, err : %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	resp := model.CreateTransactionResponse{
		OrderID:    orderID,
		TotalPrice: totalPrice,
	}

	return http.StatusOK, &model.BaseResponse{ResultData: resp}
}

// GetDetail returns the detail of a transaction by the order ID from the database,
// and the total price amount of the transaction.
func (s *transactionServiceImpl) GetDetail(ctx context.Context, orderID string) (int, *model.BaseResponse) {
	// validate request
	if strings.TrimSpace(orderID) == "" {
		return utils.RequestRequired("id")
	}

	log := logger.GetLoggerContext(ctx, "service", "GetDetail")

	transaction, err := s.transactionRepo.GetDetail(orderID)
	if err != nil {
		log.Error(fmt.Sprintf("failed to get transaction detail, err : %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	if transaction == nil {
		return http.StatusNotFound, &model.BaseResponse{ResultData: nil}
	}

	var totalAmount float64
	items := make([]model.TransactionItem, 0)

	for _, item := range transaction {
		totalAmount += item.Subtotal
		items = append(items, model.TransactionItem{
			SKU:      item.SKU,
			Quantity: item.Quantity,
			Subtotal: item.Subtotal,
		})
	}

	resp := model.GetTranscationDetailResponse{
		OrderID:     orderID,
		Items:       items,
		TotalAmount: totalAmount,
	}

	return http.StatusOK, &model.BaseResponse{ResultData: resp}
}
