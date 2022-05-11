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

// BrandService manage logical syntax for brand.
type BrandService interface {
	Create(ctx context.Context, request model.CreateBrandRequest) (int, *model.BaseResponse)
}

type brandServiceImpl struct {
	brandRepo repository.BrandRepository
}

// NewBrandService returns new instance of brandServiceImpl.
func NewBrandService() *brandServiceImpl {
	return &brandServiceImpl{}
}

// SetBrandRepo injects brand's repo for brandServiceImpl.
func (s *brandServiceImpl) SetBrandRepo(repo repository.BrandRepository) *brandServiceImpl {
	s.brandRepo = repo
	return s
}

// Validate validates if all dependency for brandServiceImpl is complete.
func (s *brandServiceImpl) Validate() *brandServiceImpl {
	if s.brandRepo == nil {
		log.Panic("Brand service need brand repository")
	}
	return s
}

// Create creates a new brand and store it into the database.
func (s *brandServiceImpl) Create(ctx context.Context, request model.CreateBrandRequest) (int, *model.BaseResponse) {
	// validate request
	if strings.TrimSpace(request.Name) == "" {
		return utils.RequestRequired("name")
	}

	log := logger.GetLoggerContext(ctx, "service", "Create")

	id, err := s.brandRepo.Create(request.Name)
	if err != nil {
		log.Error(fmt.Sprintf("failed to create brand, err: %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	resp := &model.CreateBrandResponse{
		ID: id,
	}

	return http.StatusOK, &model.BaseResponse{ResultData: resp}
}
