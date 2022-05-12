package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/project/notif-project/domain/model"
	"github.com/project/notif-project/pkg/logger"
	"github.com/project/notif-project/service"
)

// NotifHandler defines dependencies for notif handler.
type NotifHandler struct {
	notifService service.NotifService
}

// NewNotifHandler returns new instance of NotifHandler
func NewNotifHandler() *NotifHandler {
	return &NotifHandler{}
}

// SetNotifService injects notif's service for NotifHandler
func (h *NotifHandler) SetNotifService(service service.NotifService) *NotifHandler {
	h.notifService = service
	return h
}

// Validate validates if all dependency for NotifHandler is complete.
func (h *NotifHandler) Validate() *NotifHandler {
	if h.notifService == nil {
		log.Panic("Notif handler need notif service")
	}
	return h
}

// Brand handles endpoint with prefix /brand.
func (h *NotifHandler) GenerateKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLoggerContext(ctx, "handler", "Notif")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))
	log.Info(fmt.Sprintf("%+v", r))

	w.Header().Set("Content-Type", "application/json")

	var httpCode int
	var resp interface{}

	if r.Method == http.MethodPost {
		var request model.GenerateKeyRequest
		json.Unmarshal(body, &request)

		httpCode, resp = h.notifService.GenerateKey(ctx, request)
	} else {
		httpCode = http.StatusMethodNotAllowed
	}

	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(resp)
}

func (h *NotifHandler) InsertUrl(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLoggerContext(ctx, "handler", "Notif")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))
	log.Info(fmt.Sprintf("%+v", r))

	w.Header().Set("Content-Type", "application/json")

	var httpCode int
	var resp interface{}

	if r.Method == http.MethodPost {
		var request model.Url
		json.Unmarshal(body, &request)

		httpCode, resp = h.notifService.InsertUrl(ctx, request)
	} else {
		httpCode = http.StatusMethodNotAllowed
	}

	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(resp)
}

func (h *NotifHandler) SendNotificationTester(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLoggerContext(ctx, "handler", "Notif")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))
	log.Info(fmt.Sprintf("%+v", r))

	w.Header().Set("Content-Type", "application/json")

	var httpCode int
	var resp interface{}

	if r.Method == http.MethodPost {
		var request model.NotificationTesterRequest
		json.Unmarshal(body, &request)

		httpCode, resp = h.notifService.SendNotificationTester(ctx, request)
	} else {
		httpCode = http.StatusMethodNotAllowed
	}

	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(resp)
}

func (h *NotifHandler) UrlToggle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLoggerContext(ctx, "handler", "Notif")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))
	log.Info(fmt.Sprintf("%+v", r))

	w.Header().Set("Content-Type", "application/json")

	var httpCode int
	var resp interface{}

	if r.Method == http.MethodPost {
		var request model.UrlToggleStatusRequest
		json.Unmarshal(body, &request)

		httpCode, resp = h.notifService.UrlToggleStatus(ctx, request)
	} else {
		httpCode = http.StatusMethodNotAllowed
	}

	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(resp)
}

func (h *NotifHandler) SendNotification(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLoggerContext(ctx, "handler", "Notif")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))
	log.Info(fmt.Sprintf("%+v", r))

	w.Header().Set("Content-Type", "application/json")

	var httpCode int
	var resp interface{}

	if r.Method == http.MethodPost {
		var request model.SendNotif
		json.Unmarshal(body, &request)

		httpCode, resp = h.notifService.SendNotif(ctx, request)
	} else {
		httpCode = http.StatusMethodNotAllowed
	}

	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(resp)
}
