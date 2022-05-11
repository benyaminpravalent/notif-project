package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/project/notif-project/domain/model"
	"github.com/project/notif-project/domain/repository"
	"github.com/project/notif-project/pkg/logger"
	"github.com/project/notif-project/pkg/utils"
)

// NotifService manage logical syntax for notif.
type NotifService interface {
	Create(ctx context.Context, request model.GenerateKeyRequest) (int, *model.BaseResponse)
}

type notifServiceImpl struct {
	notifRepo repository.NotifRepository
}

// NewNotifService returns new instance of notifServiceImpl.
func NewNotifService() *notifServiceImpl {
	return &notifServiceImpl{}
}

// SetNotifRepo injects notif's repo for notifServiceImpl.
func (s *notifServiceImpl) SetNotifRepo(repo repository.NotifRepository) *notifServiceImpl {
	s.notifRepo = repo
	return s
}

func (s *notifServiceImpl) Validate() *notifServiceImpl {
	if s.notifRepo == nil {
		log.Panic("Notif service need notif repository")
	}
	return s
}

// Create creates a new key and store it into the database.
func (s *notifServiceImpl) Create(ctx context.Context, request model.GenerateKeyRequest) (int, *model.BaseResponse) {
	// validate request
	if request.MerchantID <= 0 {
		return utils.RequestRequired("merchant_id")
	}

	log := logger.GetLoggerContext(ctx, "service", "Create")

	uuid := uuid.New()

	finalKey := fmt.Sprintf("%d%s%d", request.MerchantID, uuid.String(), time.Now().Unix())

	err := s.notifRepo.Create(request.MerchantID, finalKey)
	if err != nil {
		log.Error(fmt.Sprintf("failed to create key err: %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	resp := &model.GenerateKeyResponse{
		Key: finalKey,
	}

	return http.StatusOK, &model.BaseResponse{ResultData: resp}
}
