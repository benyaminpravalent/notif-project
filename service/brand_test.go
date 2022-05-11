package service_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/richardsahvic/jamtangan/domain/model"
	"github.com/richardsahvic/jamtangan/service"
	"github.com/stretchr/testify/assert"

	// serviceMock "github.com/richardsahvic/jamtangan/service/mocks"
	repoMock "github.com/richardsahvic/jamtangan/domain/repository/mocks"
)

func TestCreateBrand(t *testing.T) {
	prepare()

	// TestCreateBrandEmptyRequest
	func(t *testing.T) {
		brandService := service.NewBrandService()

		// Case: empty request
		req := model.CreateBrandRequest{
			Name: "",
		}
		httpCode, resp := brandService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.Nil(t, resp.ResultData, "Result should be nil")
	}(t)

	// TestCreateBrandInternalError
	func(t *testing.T) {
		mockBrandRepo := new(repoMock.BrandRepository)
		brandService := service.NewBrandService().SetBrandRepo(mockBrandRepo)

		// Case: internal error
		req := model.CreateBrandRequest{
			Name: "jam",
		}
		mockBrandRepo.On("Create", req.Name).Return(int64(0), errors.New("error"))
		httpCode, resp := brandService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusInternalServerError)
		assert.NotNil(t, resp.RawMessage, "Response raw message should not be nil")
		mockBrandRepo.AssertNumberOfCalls(t, "Create", 1)
	}(t)

	// TestCreateBrandSuccess
	func(t *testing.T) {
		mockBrandRepo := new(repoMock.BrandRepository)
		brandService := service.NewBrandService().SetBrandRepo(mockBrandRepo)

		// Case: Success
		req := model.CreateBrandRequest{
			Name: "jam",
		}
		mockBrandRepo.On("Create", req.Name).Return(int64(1), nil)
		httpCode, resp := brandService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusOK)
		assert.NotNil(t, resp.ResultData, "Result should not be nil")
		mockBrandRepo.AssertNumberOfCalls(t, "Create", 1)
	}(t)
}
