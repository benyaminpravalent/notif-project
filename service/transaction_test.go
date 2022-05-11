package service_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/project/notif-project/domain/model"
	repoMock "github.com/project/notif-project/domain/repository/mocks"
	"github.com/project/notif-project/service"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	prepare()

	// TestCreateTransactionEmptyRequest
	func(t *testing.T) {
		transactionService := service.NewTransactionService()

		req := model.CreateTransactionRequest{
			Items: []model.TransactionItem{},
		}
		httpCode, resp := transactionService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.Nil(t, resp.ResultData)
	}(t)

	// TestCreateTransactionErrorDatabase
	// func(t *testing.T) {
	// 	mockTransactionRepo := new(repoMock.TransactionRepository)
	// 	transactionService := service.NewTransactionService().
	// 		SetTransactionRepo(mockTransactionRepo)

	// 	req := model.CreateTransactionRequest{
	// 		Items: []model.TransactionItem{
	// 			{
	// 				SKU:      "sku-test",
	// 				Quantity: 1,
	// 				Subtotal: 10000,
	// 			},
	// 		},
	// 	}
	// 	mockTransactionRepo.On("InsertList", []model.Transaction{
	// 		{
	// 			OrderID:  utils.GenerateOrderID(),
	// 			SKU:      "sku-test",
	// 			Quantity: 1,
	// 			Subtotal: 10000,
	// 		},
	// 	}).Return(errors.New("error"))
	// 	httpCode, resp := transactionService.Create(context.Background(), req)
	// 	assert.Equal(t, httpCode, http.StatusInternalServerError)
	// 	assert.Nil(t, resp.ResultData)
	// 	assert.NotEmpty(t, resp.RawMessage)
	// }(t)
}

func TestGetTransaction(t *testing.T) {
	prepare()

	// TestGetTransactionEmptyRequest
	func(t *testing.T) {
		transactionService := service.NewTransactionService()

		req := "  "
		httpCode, resp := transactionService.GetDetail(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.Nil(t, resp.ResultData)
	}(t)

	// TestGetTransactionErrorDatabase
	func(t *testing.T) {
		mockTransactionRepo := new(repoMock.TransactionRepository)
		transactionService := service.NewTransactionService().
			SetTransactionRepo(mockTransactionRepo)

		req := "orderID"
		mockTransactionRepo.On("GetDetail", req).Return(nil, errors.New("error"))
		httpCode, resp := transactionService.GetDetail(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusInternalServerError)
		assert.NotEmpty(t, resp.RawMessage)
		assert.Nil(t, resp.ResultData)
		mockTransactionRepo.AssertNumberOfCalls(t, "GetDetail", 1)
	}(t)

	// TestGetTransactionNotFound
	func(t *testing.T) {
		mockTransactionRepo := new(repoMock.TransactionRepository)
		transactionService := service.NewTransactionService().
			SetTransactionRepo(mockTransactionRepo)

		req := "orderID"
		mockTransactionRepo.On("GetDetail", req).Return(nil, nil)
		httpCode, resp := transactionService.GetDetail(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusNotFound)
		assert.Nil(t, resp.ResultData)
		assert.Empty(t, resp.RawMessage)
		mockTransactionRepo.AssertNumberOfCalls(t, "GetDetail", 1)
	}(t)

	// TestGetTransactionSuccess
	func(t *testing.T) {
		mockTransactionRepo := new(repoMock.TransactionRepository)
		transactionService := service.NewTransactionService().
			SetTransactionRepo(mockTransactionRepo)

		req := "orderID"
		mockTransactionRepo.On("GetDetail", req).Return([]*model.Transaction{}, nil)
		httpCode, resp := transactionService.GetDetail(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusOK)
		assert.NotNil(t, resp.ResultData)
		assert.Empty(t, resp.RawMessage)
		mockTransactionRepo.AssertNumberOfCalls(t, "GetDetail", 1)
	}(t)
}
