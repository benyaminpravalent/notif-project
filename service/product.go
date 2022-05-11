package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/project/notif-project/domain/model"
	"github.com/project/notif-project/domain/repository"
	"github.com/project/notif-project/pkg/logger"
	"github.com/project/notif-project/pkg/utils"
)

// ProductService manage logical syntax for product.
type ProductService interface {
	Create(ctx context.Context, request model.CreateProductRequest) (int, *model.BaseResponse)
	GetByID(ctx context.Context, productID string) (int, *model.BaseResponse)
	GetByBrandID(ctx context.Context, brandID string) (int, *model.BaseResponse)
}

type productServiceImpl struct {
	productRepo repository.ProductRepository
	brandRepo   repository.BrandRepository
}

// NewProductService returns new instance of productServiceImpl.
func NewProductService() *productServiceImpl {
	return &productServiceImpl{}
}

// SetProductRepo injects product's repo for productServiceImpl.
func (s *productServiceImpl) SetProductRepo(repo repository.ProductRepository) *productServiceImpl {
	s.productRepo = repo
	return s
}

// SetBrandRepo injects brand's repo for productServiceImpl.
func (s *productServiceImpl) SetBrandRepo(repo repository.BrandRepository) *productServiceImpl {
	s.brandRepo = repo
	return s
}

// Validate validates if all dependency for productServiceImpl is complete.
func (s *productServiceImpl) Validate() *productServiceImpl {
	if s.productRepo == nil {
		log.Panic("Product service need product repository")
	}
	if s.brandRepo == nil {
		log.Panic("Product service need brand repository")
	}
	return s
}

// Create creates a new product and store it into the database.
func (s *productServiceImpl) Create(ctx context.Context, request model.CreateProductRequest) (int, *model.BaseResponse) {
	// validate request
	if request.BrandID == 0 {
		return utils.RequestRequired("brand_id")
	} else if strings.TrimSpace(request.SKU) == "" {
		return utils.RequestRequired("sku")
	} else if request.Price == 0 {
		return utils.RequestRequired("price")
	}

	log := logger.GetLoggerContext(ctx, "service", "Create")

	brand, err := s.brandRepo.GetByID(request.BrandID)
	if err != nil {
		log.Error(fmt.Sprintf("failed to get brand, err : %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	if brand == nil {
		return utils.RequestInvalid("brand_id")
	}

	checkProduct, err := s.productRepo.GetBySKU(request.SKU)
	if err != nil {
		log.Error(fmt.Sprintf("failed to get product by SKU, err : %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	if checkProduct != nil {
		return utils.RequestInvalid("sku")
	}

	product := model.Product{
		BrandID: request.BrandID,
		SKU:     request.SKU,
		Stock:   request.Stock,
		Price:   request.Price,
	}

	err = s.productRepo.Create(&product)
	if err != nil {
		log.Error(fmt.Sprintf("failed to create product, err : %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	resp := &model.CreateProductResponse{
		ID: product.ID,
	}

	return http.StatusOK, &model.BaseResponse{ResultData: resp}
}

// GetByID returns a product details by the ID from the database .
func (s *productServiceImpl) GetByID(ctx context.Context, productID string) (int, *model.BaseResponse) {
	// validate request
	if strings.TrimSpace(productID) == "" {
		return utils.RequestRequired("id")
	}

	id, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return utils.RequestInvalid("id")
	}

	log := logger.GetLoggerContext(ctx, "service", "GetByID")

	product, err := s.productRepo.GetByID(id)
	if err != nil {
		log.Error(fmt.Sprintf("failed to get product by id, err : %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	if product == nil {
		return http.StatusNotFound, &model.BaseResponse{}
	}

	productResp := model.GetProductResponse{
		ID:      product.ID,
		BrandID: product.BrandID,
		SKU:     product.SKU,
		Stock:   product.Stock,
		Price:   product.Price,
	}

	return http.StatusOK, &model.BaseResponse{ResultData: productResp}
}

// GetByBrandID returns a list of product by the brand's ID from the database.
func (s *productServiceImpl) GetByBrandID(ctx context.Context, brandID string) (int, *model.BaseResponse) {
	// validate request
	if strings.TrimSpace(brandID) == "" {
		return utils.RequestRequired("id")
	}

	id, err := strconv.ParseInt(brandID, 10, 64)
	if err != nil {
		return utils.RequestInvalid("id")
	}

	log := logger.GetLoggerContext(ctx, "service", "GetByBrandID")

	product, err := s.productRepo.GetByBrandID(id)
	if err != nil {
		log.Error(fmt.Sprintf("failed to get product by brand id, err : %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	resp := model.GetProductByBrandIDResponse{
		Products: product,
	}

	return http.StatusOK, &model.BaseResponse{ResultData: resp}
}
