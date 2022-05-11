package service_test

import (
	"context"
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

		httpCode, resp := notifService.Create(context.Background(), req)
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

		// mockNotifRepo.On("Create", req.MerchantID).Return(errors.New("error"))
		httpCode, resp := notifService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.NotNil(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "Create", 0)
	}(t)

	// Test Create Notif Success
	func(t *testing.T) {
		mockNotifRepo := new(repoMock.NotifRepository)
		notifService := service.NewNotifService().SetNotifRepo(mockNotifRepo)

		// Case : Success
		req := model.GenerateKeyRequest{
			MerchantID: 1,
		}

		mockNotifRepo.On("Create", req.MerchantID, mock.Anything).Return(nil)
		httpCode, resp := notifService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusOK)
		assert.NotNil(t, resp.RawMessage, "Response raw message should not be nil")
		mockNotifRepo.AssertNumberOfCalls(t, "Create", 1)
	}(t)
}
