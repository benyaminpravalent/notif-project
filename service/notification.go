package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	GenerateKey(ctx context.Context, request model.GenerateKeyRequest) (int, *model.BaseResponse)
	InsertUrl(ctx context.Context, request model.Url) (int, *model.BaseResponse)
	NotificationTester(ctx context.Context, request model.NotificationTesterRequest) (int, *model.BaseResponse)
	UrlToggleStatus(ctx context.Context, request model.UrlToggleStatusRequest) (int, *model.BaseResponse)
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
func (s *notifServiceImpl) GenerateKey(ctx context.Context, request model.GenerateKeyRequest) (int, *model.BaseResponse) {
	// validate request
	if request.MerchantID <= 0 {
		return utils.RequestRequired("merchant_id")
	}

	log := logger.GetLoggerContext(ctx, "service", "GenerateKey")

	uuid := uuid.New()

	finalKey := fmt.Sprintf("%d%s%d", request.MerchantID, uuid.String(), time.Now().Unix())

	err := s.notifRepo.GenerateKey(request.MerchantID, finalKey)
	if err != nil {
		log.Error(fmt.Sprintf("failed to create key err: %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	resp := &model.GenerateKeyResponse{
		Key: finalKey,
	}

	return http.StatusOK, &model.BaseResponse{ResultData: resp}
}

func (s *notifServiceImpl) InsertUrl(ctx context.Context, request model.Url) (int, *model.BaseResponse) {
	// validate request
	if request.MerchantID <= 0 {
		return utils.RequestRequired("merchant_id")
	} else if request.Url == "" {
		return utils.RequestRequired("url")
	} else if request.NotificationType == "" {
		return utils.RequestRequired("notification_type")
	}

	log := logger.GetLoggerContext(ctx, "service", "InsertUrl")

	count, err := s.notifRepo.CheckUrlExistence(request)
	if err != nil {
		log.Error(fmt.Sprintf("failed while do CheckUrlExistence, err: %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	if count > 0 {
		log.Error(fmt.Sprintf("Merchant with ID=%d already has the url", request.MerchantID))
		return http.StatusBadRequest, &model.BaseResponse{RawMessage: errors.New("Url or the notification_type is already exist").Error()}
	}

	err = s.notifRepo.InsertUrl(request)
	if err != nil {
		log.Error(fmt.Sprintf("failed while do InsertUrl, err: %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	return http.StatusOK, &model.BaseResponse{ResultData: "URL is successfully created"}
}

func (s *notifServiceImpl) NotificationTester(ctx context.Context, request model.NotificationTesterRequest) (int, *model.BaseResponse) {
	// validate request
	if request.UrlID <= 0 {
		return utils.RequestRequired("url_id")
	}

	log := logger.GetLoggerContext(ctx, "service", "InsertUrl")

	data, err := s.notifRepo.GetUrlDetail(request.UrlID)
	if err != nil {
		log.Error(fmt.Sprintf("failed while do GetUrlDetail, err: %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	if data.IsActive {
		log.Error(fmt.Sprintf("Merchant with ID=%d, the url is active, can't test the notification", data.MerchantID))
		return http.StatusBadRequest, &model.BaseResponse{RawMessage: errors.New("URL status is Active").Error()}
	}

	go func(notification_type string) {
		postBody, _ := json.Marshal(map[string]interface{}{
			"notification_type":  notification_type,
			"transaction_id":     123,
			"amount":             100000,
			"transaction_status": "success",
		})
		responseBody := bytes.NewBuffer(postBody)

		resp, err := http.Post(data.Url, "application/json", responseBody)
		//Error handling
		if err != nil {
			log.Error(fmt.Sprintf("An Error Occured with merchantID=%d, err : %v", data.MerchantID, err))
			return
		}

		defer resp.Body.Close()

		//Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Printf(sb)
	}(data.NotificationType)

	return http.StatusOK, &model.BaseResponse{ResultData: "SUCCESS"}
}

func (s *notifServiceImpl) UrlToggleStatus(ctx context.Context, request model.UrlToggleStatusRequest) (int, *model.BaseResponse) {
	// validate request
	if request.UrlID <= 0 {
		return utils.RequestRequired("url_id")
	}

	log := logger.GetLoggerContext(ctx, "service", "UrlToggleStatus")

	err := s.notifRepo.UrlToggleStatus(request.UrlID)
	if err != nil {
		log.Error(fmt.Sprintf("failed while do UrlToggleStatus, err: %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	return http.StatusOK, &model.BaseResponse{ResultData: "Url Toggle is successfully executed"}
}
