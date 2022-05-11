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
	database.InitMySql(ctx)

	// REPOSITORIES
	brandRepo := repository.NewBrandRepository()
	productRepo := repository.NewProductRepository()
	transactionRepo := repository.NewTransactionRepository()

	brandService := service.NewBrandService().
		SetBrandRepo(brandRepo).
		Validate()

	productService := service.NewProductService().
		SetProductRepo(productRepo).
		SetBrandRepo(brandRepo).
		Validate()

	transactionService := service.NewTransactionService().
		SetTransactionRepo(transactionRepo).
		Validate()

	brandHandler := handler.NewBrandHandler().
		SetBrandService(brandService).
		Validate()

	productHandler := handler.NewProductHandler().
		SetProductService(productService).
		Validate()

	transactionHandler := handler.NewTransactionhandler().
		SetTransactionService(transactionService).
		Validate()

	route := http.NewServeMux()

	// Brand API
	route.HandleFunc("/brand", brandHandler.Brand)

	// Product API
	route.HandleFunc("/product", productHandler.Product)
	route.HandleFunc("/product/brand", productHandler.ProductByBrand)

	// Transaction API
	route.HandleFunc("/order", transactionHandler.Transaction)

	log.Println("SERVER STARTED")

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetString("port")), route)
}
