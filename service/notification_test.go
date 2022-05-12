package service_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/project/notif-project/cmd"
	"github.com/project/notif-project/domain/model"
	"github.com/project/notif-project/pkg/config"
	"github.com/project/notif-project/pkg/logger"
	"github.com/project/notif-project/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// serviceMock "github.com/project/notif-project/service/mocks"
	repoMock "github.com/project/notif-project/domain/repository/mocks"
)

func prepare() {
	config.Load(cmd.DefaultConfig, "../config.json")
	logger.Configure()
	return
}

func TestGenerateKey(t *testing.T) {
	prepare()

	// TestGenerateKeyEmptyRequest
	func(t *testing.T) {
		notifService := service.NewNotifService()

		//Case: empty request
		req := model.GenerateKeyRequest{
			MerchantID: 0,
		}

		httpCode, resp := notifService.GenerateKey(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.Nil(t, resp.ResultData, "Result should be nil")
	}(t)

	// TestGenerateKey Failed
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		// Case : internal error
		req := model.GenerateKeyRequest{
			MerchantID: -100,
		}

		// mockNotifRepo.On("GenerateKey", req.MerchantID).Return(errors.New("error"))
		httpCode, resp := notifService.GenerateKey(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.NotNil(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "GenerateKey", 0)
	}(t)

	// Test Create Notif Success
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		// Case : Success
		req := model.GenerateKeyRequest{
			MerchantID: 1,
		}

		mockNotifRepo.On("GenerateKey", req.MerchantID, mock.Anything).Return(nil)
		httpCode, resp := notifService.GenerateKey(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusOK)
		assert.NotNil(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "GenerateKey", 1)
	}(t)
}

func TestUrlToggleStatus(t *testing.T) {
	prepare()

	// Case when  merchantID = 0
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.Url{
			MerchantID: 0,
		}

		httpCode, resp := notifService.InsertUrl(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.NotNil(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "CheckUrlExistence", 0)
		mockNotifRepo.AssertNumberOfCalls(t, "InsertUrl", 0)
	}(t)

	// Case when  url = ""
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.Url{
			Url: "",
		}

		httpCode, resp := notifService.InsertUrl(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.NotNil(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "CheckUrlExistence", 0)
		mockNotifRepo.AssertNumberOfCalls(t, "InsertUrl", 0)
	}(t)

	// Case when  notification_type = ""
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.Url{
			NotificationType: "",
		}

		httpCode, resp := notifService.InsertUrl(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.NotNil(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "CheckUrlExistence", 0)
		mockNotifRepo.AssertNumberOfCalls(t, "InsertUrl", 0)
	}(t)

	// Case when  success
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.Url{
			MerchantID:       4,
			Url:              "http://localhost:3000/tokobaju/webhook/notification",
			NotificationType: "refund",
		}

		mockNotifRepo.On("CheckUrlExistence", req).Return(int64(0), nil)
		mockNotifRepo.On("InsertUrl", req).Return(nil)
		httpCode, resp := notifService.InsertUrl(context.Background(), req)
		assert.Equal(t, http.StatusOK, httpCode)
		assert.Empty(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "CheckUrlExistence", 1)
		mockNotifRepo.AssertNumberOfCalls(t, "InsertUrl", 1)
	}(t)

	// Case when  notification_type = ""
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.Url{
			MerchantID:       4,
			Url:              "http://localhost:3000/tokobaju/webhook/notification",
			NotificationType: "refund",
		}

		mockNotifRepo.On("CheckUrlExistence", req).Return(int64(1), nil)
		httpCode, resp := notifService.InsertUrl(context.Background(), req)
		assert.Equal(t, http.StatusBadRequest, httpCode)
		assert.NotEmpty(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "CheckUrlExistence", 1)
	}(t)
}

func TestInsertUrl(t *testing.T) {
	prepare()

	// Case when  urlID <= 0
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.UrlToggleStatusRequest{
			UrlID: 0,
		}

		httpCode, resp := notifService.UrlToggleStatus(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.NotEmpty(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "UrlToggleStatus", 0)
	}(t)

	// Case when setting urlID
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.UrlToggleStatusRequest{
			UrlID: 100000,
		}

		mockNotifRepo.On("UrlToggleStatus", req.UrlID).Return(errors.New("error"))
		httpCode, resp := notifService.UrlToggleStatus(context.Background(), req)
		assert.Equal(t, http.StatusInternalServerError, httpCode)
		assert.NotEmpty(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "UrlToggleStatus", 1)
	}(t)

	// Case when setting urlID success
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.UrlToggleStatusRequest{
			UrlID: 4,
		}

		mockNotifRepo.On("UrlToggleStatus", req.UrlID).Return(nil)
		httpCode, resp := notifService.UrlToggleStatus(context.Background(), req)
		assert.Equal(t, http.StatusOK, httpCode)
		assert.Empty(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "UrlToggleStatus", 1)
	}(t)
}

func TestSendNotificationTester(t *testing.T) {
	prepare()

	// Case when  urlID <= 0
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.UrlToggleStatusRequest{
			UrlID: 0,
		}

		httpCode, resp := notifService.UrlToggleStatus(context.Background(), req)
		assert.Equal(t, http.StatusBadRequest, httpCode)
		assert.NotEmpty(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "GetUrlDetail", 0)
	}(t)

	// Case when urlID is bad request
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.NotificationTesterRequest{
			UrlID: 4,
		}

		mockNotifRepo.On("GetUrlDetail", req.UrlID).Return(model.GetUrlDetailRes{IsActive: true}, nil)
		httpCode, resp := notifService.SendNotificationTester(context.Background(), req)
		assert.Equal(t, http.StatusBadRequest, httpCode)
		assert.NotEmpty(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "GetUrlDetail", 1)
	}(t)

	// Case when urlID is bad request
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.NotificationTesterRequest{
			UrlID: 4,
		}

		mockNotifRepo.On("GetUrlDetail", req.UrlID).Return(model.GetUrlDetailRes{IsActive: false}, nil)
		httpCode, resp := notifService.SendNotificationTester(context.Background(), req)
		assert.Equal(t, http.StatusOK, httpCode)
		assert.Empty(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "GetUrlDetail", 1)
	}(t)
}

func TestRetrySendNotifExecution(t *testing.T) {
	prepare()

	// Case when failed
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.SendNotifGoRoutine{
			MerchantID:        4,
			NotificationType:  "payment",
			TransactionID:     2323,
			Amount:            10000.00,
			TransactionStatus: "success",
			CheckSum:          "dkjgfndskjgnds9u3t893uytdsf8eyt8ghgidss",
			IdempotencyKey:    "f43t3409it29fdid385y985398534jnbfjbndf",
			Url:               "http://localhost:3000/tokobaju/webhook/notification",
		}

		reqRepo := model.UpdateNotifStatus{
			MerchantID:         req.MerchantID,
			Key:                req.IdempotencyKey,
			TransactionID:      req.TransactionID,
			NotificationStatus: "success",
		}

		mockNotifRepo.On("UpdateNotifStatus", reqRepo).Return(errors.New("error"))
		httpCode, resp := notifService.RetrySendNotifExecution(context.Background(), req)
		assert.Equal(t, http.StatusInternalServerError, httpCode)
		assert.NotEmpty(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "UpdateNotifStatus", 1)
	}(t)

	// Case when success
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		req := model.SendNotifGoRoutine{
			MerchantID:        4,
			NotificationType:  "payment",
			TransactionID:     2323,
			Amount:            10000.00,
			TransactionStatus: "success",
			CheckSum:          "dkjgfndskjgnds9u3t893uytdsf8eyt8ghgidss",
			IdempotencyKey:    "f43t3409it29fdid385y985398534jnbfjbndf",
			Url:               "http://localhost:3000/tokobaju/webhook/notification",
		}

		reqRepo := model.UpdateNotifStatus{
			MerchantID:         req.MerchantID,
			Key:                req.IdempotencyKey,
			TransactionID:      req.TransactionID,
			NotificationStatus: "success",
		}

		mockNotifRepo.On("UpdateNotifStatus", reqRepo).Return(nil)
		httpCode, resp := notifService.RetrySendNotifExecution(context.Background(), req)
		assert.Equal(t, http.StatusOK, httpCode)
		assert.Empty(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "UpdateNotifStatus", 1)
	}(t)
}
