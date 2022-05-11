package service_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/project/notif-project/cmd"
	"github.com/project/notif-project/domain/model"
	repoMock "github.com/project/notif-project/domain/repository/mocks"
	"github.com/project/notif-project/pkg/config"
	"github.com/project/notif-project/pkg/logger"
	"github.com/project/notif-project/service"

	"github.com/stretchr/testify/assert"
)

func prepare() {
	config.Load(cmd.DefaultConfig, "../config.json")
	logger.Configure()
	return
}

func TestCreateProduct(t *testing.T) {
	prepare()

	// TestCreateProductEmptyRequest
	func(t *testing.T) {
		productService := service.NewProductService()

		// Case: empty brand ID
		req := model.CreateProductRequest{
			BrandID: 0,
		}
		httpCode, resp := productService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.Nil(t, resp.ResultData)

		// Case: empty SKU
		req = model.CreateProductRequest{
			SKU: "   ",
		}
		httpCode, resp = productService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.Nil(t, resp.ResultData)

		// Case: empty price
		req = model.CreateProductRequest{
			Price: 0,
		}
		httpCode, resp = productService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.Nil(t, resp.ResultData)
	}(t)

	// TestCreateProductInvalidBrandCheck
	func(t *testing.T) {
		mockBrandRepo := new(repoMock.BrandRepository)
		productService := service.NewProductService().SetBrandRepo(mockBrandRepo)

		// Case: invalid brand ID
		req := model.CreateProductRequest{
			BrandID: 1,
			SKU:     "sku-test",
			Price:   100,
		}
		mockBrandRepo.On("GetByID", req.BrandID).Return(nil, nil)
		httpCode, resp := productService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.NotEmpty(t, resp.RawMessage)
		mockBrandRepo.AssertNumberOfCalls(t, "GetByID", 1)
	}(t)

	// TestCreateProductInvalidSKUCheck
	func(t *testing.T) {
		mockBrandRepo := new(repoMock.BrandRepository)
		mockProductRepo := new(repoMock.ProductRepository)
		productService := service.NewProductService().
			SetBrandRepo(mockBrandRepo).
			SetProductRepo(mockProductRepo)

		// Case: invalid brand ID
		req := model.CreateProductRequest{
			BrandID: 1,
			SKU:     "sku-test",
			Price:   100,
		}
		mockBrandRepo.On("GetByID", req.BrandID).Return(&model.Brand{ID: 1}, nil)
		mockProductRepo.On("GetBySKU", req.SKU).Return(&model.Product{SKU: req.SKU}, nil)
		httpCode, resp := productService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.NotEmpty(t, resp.RawMessage)
		mockBrandRepo.AssertNumberOfCalls(t, "GetByID", 1)
		mockProductRepo.AssertNumberOfCalls(t, "GetBySKU", 1)
	}(t)

	// TestCreateProductFailedCreateCheck
	func(t *testing.T) {
		mockBrandRepo := new(repoMock.BrandRepository)
		mockProductRepo := new(repoMock.ProductRepository)
		productService := service.NewProductService().
			SetBrandRepo(mockBrandRepo).
			SetProductRepo(mockProductRepo)

		// Case: invalid brand ID
		req := model.CreateProductRequest{
			BrandID: 1,
			SKU:     "sku-test",
			Price:   100,
		}
		result := &model.Product{
			BrandID: req.BrandID,
			SKU:     req.SKU,
			Price:   req.Price,
		}
		mockBrandRepo.On("GetByID", req.BrandID).Return(&model.Brand{ID: 1}, nil)
		mockProductRepo.On("GetBySKU", req.SKU).Return(nil, nil)
		mockProductRepo.On("Create", result).Return(errors.New("error"))
		httpCode, resp := productService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusInternalServerError)
		assert.NotEmpty(t, resp.RawMessage)
		mockBrandRepo.AssertNumberOfCalls(t, "GetByID", 1)
		mockProductRepo.AssertNumberOfCalls(t, "GetBySKU", 1)
		mockProductRepo.AssertNumberOfCalls(t, "Create", 1)
	}(t)

	// TestCreateProductSuccess
	func(t *testing.T) {
		mockBrandRepo := new(repoMock.BrandRepository)
		mockProductRepo := new(repoMock.ProductRepository)
		productService := service.NewProductService().
			SetBrandRepo(mockBrandRepo).
			SetProductRepo(mockProductRepo)

		// Case: invalid brand ID
		req := model.CreateProductRequest{
			BrandID: 1,
			SKU:     "sku-test",
			Price:   100,
		}
		result := &model.Product{
			BrandID: req.BrandID,
			SKU:     req.SKU,
			Price:   req.Price,
		}
		mockBrandRepo.On("GetByID", req.BrandID).Return(&model.Brand{ID: 1}, nil)
		mockProductRepo.On("GetBySKU", req.SKU).Return(nil, nil)
		mockProductRepo.On("Create", result).Return(nil)
		httpCode, resp := productService.Create(context.Background(), req)
		assert.Equal(t, httpCode, http.StatusOK)
		assert.Empty(t, resp.RawMessage)
		mockBrandRepo.AssertNumberOfCalls(t, "GetByID", 1)
		mockProductRepo.AssertNumberOfCalls(t, "GetBySKU", 1)
		mockProductRepo.AssertNumberOfCalls(t, "Create", 1)
	}(t)
}

func TestGetByID(t *testing.T) {
	prepare()

	// TestGetByIDRequest
	func(t *testing.T) {
		productService := service.NewProductService()

		// Case: empty product ID
		id := " "
		httpCode, resp := productService.GetByID(context.Background(), id)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.Nil(t, resp.ResultData)

		// Case: invalid type product ID
		id = "a"
		httpCode, resp = productService.GetByID(context.Background(), id)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.Nil(t, resp.ResultData)
	}(t)

	// TestGetByIDErrorDatabase
	func(t *testing.T) {
		mockProductRepo := new(repoMock.ProductRepository)
		productService := service.NewProductService().SetProductRepo(mockProductRepo)

		id := "1"
		mockProductRepo.On("GetByID", int64(1)).Return(nil, errors.New("error"))
		httpCode, resp := productService.GetByID(context.Background(), id)
		assert.Equal(t, httpCode, http.StatusInternalServerError)
		assert.NotEmpty(t, resp.RawMessage)
		mockProductRepo.AssertNumberOfCalls(t, "GetByID", 1)
	}(t)

	// TestGetByIDDataNotFound
	func(t *testing.T) {
		mockProductRepo := new(repoMock.ProductRepository)
		productService := service.NewProductService().SetProductRepo(mockProductRepo)

		id := "1"
		mockProductRepo.On("GetByID", int64(1)).Return(nil, nil)
		httpCode, resp := productService.GetByID(context.Background(), id)
		assert.Equal(t, httpCode, http.StatusNotFound)
		assert.Empty(t, resp.RawMessage)
		mockProductRepo.AssertNumberOfCalls(t, "GetByID", 1)
	}(t)

	// TestGetByIDSuccess
	func(t *testing.T) {
		mockProductRepo := new(repoMock.ProductRepository)
		productService := service.NewProductService().SetProductRepo(mockProductRepo)

		id := "1"
		mockProductRepo.On("GetByID", int64(1)).Return(&model.Product{
			ID: 1,
		}, nil)
		httpCode, resp := productService.GetByID(context.Background(), id)
		assert.Equal(t, httpCode, http.StatusOK)
		assert.Empty(t, resp.RawMessage)
		assert.NotNil(t, resp.ResultData)
		mockProductRepo.AssertNumberOfCalls(t, "GetByID", 1)
	}(t)
}

func TestGetByBrandID(t *testing.T) {
	prepare()

	// TestGetByBrandIDEmptyRequest
	func(t *testing.T) {
		productService := service.NewProductService()

		// Case: empty brand ID
		id := " "
		httpCode, resp := productService.GetByBrandID(context.Background(), id)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.Nil(t, resp.ResultData)

		// Case: invalid type brand ID
		id = "a"
		httpCode, resp = productService.GetByID(context.Background(), id)
		assert.Equal(t, httpCode, http.StatusBadRequest)
		assert.Nil(t, resp.ResultData)
	}(t)

	// TestGetByBrandIDErrorDatabase
	func(t *testing.T) {
		mockProductRepo := new(repoMock.ProductRepository)
		productService := service.NewProductService().SetProductRepo(mockProductRepo)

		id := "1"
		mockProductRepo.On("GetByBrandID", int64(1)).Return(nil, errors.New("error"))
		httpCode, resp := productService.GetByBrandID(context.Background(), id)
		assert.Equal(t, httpCode, http.StatusInternalServerError)
		assert.NotEmpty(t, resp.RawMessage)
		mockProductRepo.AssertNumberOfCalls(t, "GetByBrandID", 1)
	}(t)

	// TestGetByBrandIDSuccess
	func(t *testing.T) {
		mockProductRepo := new(repoMock.ProductRepository)
		productService := service.NewProductService().SetProductRepo(mockProductRepo)

		id := "1"
		mockProductRepo.On("GetByBrandID", int64(1)).Return([]*model.Product{}, nil)
		httpCode, resp := productService.GetByBrandID(context.Background(), id)
		assert.Equal(t, httpCode, http.StatusOK)
		assert.Empty(t, resp.RawMessage)
		assert.NotEmpty(t, resp.ResultData)
		mockProductRepo.AssertNumberOfCalls(t, "GetByBrandID", 1)
	}(t)
}
