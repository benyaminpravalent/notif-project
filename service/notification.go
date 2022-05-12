package service

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
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
	GenerateKey(ctx context.Context, request model.GenerateKeyRequest) (int, *model.BaseResponse)
	InsertUrl(ctx context.Context, request model.Url) (int, *model.BaseResponse)
	SendNotificationTester(ctx context.Context, request model.NotificationTesterRequest) (int, *model.BaseResponse)
	UrlToggleStatus(ctx context.Context, request model.UrlToggleStatusRequest) (int, *model.BaseResponse)
	SendNotif(ctx context.Context, request model.SendNotif) (int, *model.BaseResponse)
	SendNotifExecution(ctx context.Context, param model.SendNotifGoRoutine) (int, *model.BaseResponse)
	RetrySendNotifExecution(ctx context.Context, param model.SendNotifGoRoutine) (int, *model.BaseResponse)
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

func (s *notifServiceImpl) SendNotificationTester(ctx context.Context, request model.NotificationTesterRequest) (int, *model.BaseResponse) {
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
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Error(err.Error())
		// }
		// sb := string(body)
		// log.Printf(sb)
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

func (s *notifServiceImpl) SendNotif(ctx context.Context, request model.SendNotif) (int, *model.BaseResponse) {
	// validate request
	if request.TransactionID <= 0 {
		return utils.RequestRequired("transaction_id")
	} else if request.MerchantID <= 0 {
		return utils.RequestRequired("merchant_id")
	} else if request.NotificationType == "" {
		return utils.RequestRequired("notification_type")
	} else if request.TransactionStatus == "" {
		return utils.RequestRequired("transaction_status")
	}

	log := logger.GetLoggerContext(ctx, "service", "SendNotif")

	// get the detail of url for validation
	data, err := s.notifRepo.GetMerchantUrlDetail(request.MerchantID, request.NotificationType)
	if err != nil {
		log.Error(fmt.Sprintf("failed while do GetMerchantUrlDetail, err: %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	if !data.IsActive {
		log.Error(fmt.Sprintf("Merchant with ID=%d, the url is Inactive, can't test the notification", request.MerchantID))
		return http.StatusBadRequest, &model.BaseResponse{RawMessage: errors.New("URL status is InActive").Error()}
	}

	checkOnProsessNotifiParam := model.CheckOnProsessNotif{
		MerchantID:    request.MerchantID,
		TransactionID: request.TransactionID,
	}

	// check if there is already notification with current merchantID and transactionID
	count, err := s.notifRepo.CheckOnProsessNotif(checkOnProsessNotifiParam)
	if err != nil {
		log.Error(fmt.Sprintf("failed while do CheckOnProsessNotif, err: %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	if count > 0 {
		log.Error(fmt.Sprintf("Merchant with ID=%d, there's already notification with transactionID= %d", request.MerchantID, request.TransactionID))
		return http.StatusBadRequest, &model.BaseResponse{RawMessage: errors.New("Url or the notification_type is already exist").Error()}
	}

	uuid := uuid.New()
	idempotencyKey := fmt.Sprintf("%s%d", uuid.String(), time.Now().Unix())

	// record every notification that able to execute
	err = s.notifRepo.InsertNotifExecution(model.InsertNotifExecution{
		MerchantID:        request.MerchantID,
		UrlID:             data.UrlID,
		NotificationType:  request.NotificationType,
		TransactionID:     request.TransactionID,
		Amount:            request.Amount,
		TransactionStatus: request.TransactionStatus,
		Key:               idempotencyKey,
	})
	if err != nil {
		log.Error(fmt.Sprintf("failed while do InsertNotifExecution, err: %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	// create checksum to authenticate the notif is from Xendit
	temp := fmt.Sprintf("%s%d%.2f%s%s", request.NotificationType, request.TransactionID, request.Amount, request.TransactionStatus, data.MerchantKey)
	md5 := md5.Sum([]byte(temp))
	checkSum := fmt.Sprintf("%x", md5)

	// start to send notif through merchant's webhook
	go s.SendNotifExecution(ctx, model.SendNotifGoRoutine{
		MerchantID:        request.MerchantID,
		Url:               data.Url,
		NotificationType:  request.NotificationType,
		TransactionID:     request.TransactionID,
		Amount:            request.Amount,
		TransactionStatus: request.TransactionStatus,
		CheckSum:          checkSum,
		IdempotencyKey:    idempotencyKey,
	})

	return http.StatusOK, &model.BaseResponse{ResultData: "SUCCESS"}
}

func (s *notifServiceImpl) SendNotifExecution(ctx context.Context, param model.SendNotifGoRoutine) (int, *model.BaseResponse) {
	log := logger.GetLoggerContext(ctx, "service", "SendNotifExecution")

	postBody, _ := json.Marshal(map[string]interface{}{
		"notification_type":  param.NotificationType,
		"transaction_id":     param.TransactionID,
		"amount":             param.Amount,
		"transaction_status": param.TransactionStatus,
		"check_sum":          param.CheckSum,
		"idempotency_key":    param.IdempotencyKey,
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(param.Url, "application/json", responseBody)
	//Error handling
	if err != nil {
		log.Error(fmt.Sprintf("An Error Occured while try to HIT Webhook with merchantID=%d, err : %v", param.MerchantID, err))
		log.Error("Service Will do retry for 4 times")
		httpCode, _ := s.RetrySendNotifExecution(ctx, param)
		if httpCode != http.StatusOK {
			log.Error(fmt.Sprintf("An Error Occured while do RetrySendNotifExecution with merchantID=%d, err : %v", param.MerchantID, err))
		}
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	// update notification_status
	err = s.notifRepo.UpdateNotifStatus(model.UpdateNotifStatus{
		MerchantID:         param.MerchantID,
		Key:                param.IdempotencyKey,
		TransactionID:      param.TransactionID,
		NotificationStatus: "success",
	})
	if err != nil {
		log.Error(fmt.Sprintf("failed while do UpdateNotifStatus, err: %s", err.Error()))
		return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
	}

	defer resp.Body.Close()

	//Read the response body
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Error(err.Error())
	// }
	// sb := string(body)
	// log.Printf(sb)

	return http.StatusOK, &model.BaseResponse{ResultData: "SUCCESS"}
}

func (s *notifServiceImpl) RetrySendNotifExecution(ctx context.Context, param model.SendNotifGoRoutine) (int, *model.BaseResponse) {
	log := logger.GetLoggerContext(ctx, "service", "RetrySendNotifExecution")

	for i := 0; i < 4; i++ {
		postBody, _ := json.Marshal(map[string]interface{}{
			"notification_type":  param.NotificationType,
			"transaction_id":     param.TransactionID,
			"amount":             param.Amount,
			"transaction_status": param.TransactionStatus,
			"check_sum":          param.CheckSum,
			"idempotency_key":    param.IdempotencyKey,
		})
		responseBody := bytes.NewBuffer(postBody)

		_, err := http.Post(param.Url, "application/json", responseBody)
		//Error handling
		if err != nil {
			if i == 3 {
				// update notification_status
				err = s.notifRepo.UpdateNotifStatus(model.UpdateNotifStatus{
					MerchantID:         param.MerchantID,
					Key:                param.IdempotencyKey,
					TransactionID:      param.TransactionID,
					NotificationStatus: "failed",
				})
				if err != nil {
					log.Error(fmt.Sprintf("failed while do UpdateNotifStatus, err: %s", err.Error()))
					return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
				}
				return http.StatusOK, &model.BaseResponse{ResultData: "SUCCESS"}
			}
			time.Sleep(5 * time.Second)
			continue
		}

		// update notification_status
		err = s.notifRepo.UpdateNotifStatus(model.UpdateNotifStatus{
			MerchantID:         param.MerchantID,
			Key:                param.IdempotencyKey,
			TransactionID:      param.TransactionID,
			NotificationStatus: "success",
		})
		if err != nil {
			log.Error(fmt.Sprintf("failed while do UpdateNotifStatus, err: %s", err.Error()))
			return http.StatusInternalServerError, &model.BaseResponse{RawMessage: err.Error()}
		}

		break
	}

	return http.StatusOK, &model.BaseResponse{ResultData: "SUCCESS"}
}
