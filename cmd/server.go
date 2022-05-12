package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/project/notif-project/api/handler"
	"github.com/project/notif-project/domain/repository"
	"github.com/project/notif-project/pkg/config"
	"github.com/project/notif-project/pkg/constant"
	"github.com/project/notif-project/pkg/database"
	"github.com/project/notif-project/pkg/logger"
	"github.com/project/notif-project/service"
)

// StartServer starts the server.
func StartServer() {
	ctx := context.Background()

	if err := config.Load(DefaultConfig, constant.ConfigURL); err != nil {
		log.Fatal(err)
	}

	logger.Configure()
	database.InitDatabases(ctx)

	// REPOSITORIES
	notifRepo := repository.NewNotifRepository()

	notifService := service.NewNotifService().
		SetNotifRepo(notifRepo).
		Validate()

	notifHandler := handler.NewNotifHandler().
		SetNotifService(notifService).
		Validate()

	route := http.NewServeMux()

	// Notif API
	route.HandleFunc("/key/create", notifHandler.GenerateKey)
	route.HandleFunc("/notif/create", notifHandler.InsertUrl)
	route.HandleFunc("/notif/test", notifHandler.SendNotificationTester)
	route.HandleFunc("/notif/toggle", notifHandler.UrlToggle)
	route.HandleFunc("/notif/execute", notifHandler.SendNotification)

	log.Println(fmt.Sprintf("SERVER STARTED IN PORT : %s", config.GetString("port")))

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetString("port")), route)
}
