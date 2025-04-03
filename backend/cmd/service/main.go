package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	server "backend/internal/api"
	"backend/internal/api/delivery"
	"backend/internal/api/health"
	"backend/internal/api/list"
	"backend/internal/api/upload"
	"backend/internal/generated/api"
	delivery_usecase "backend/internal/usecase/delivery"
	list_usecase "backend/internal/usecase/list"
	upload_usecase "backend/internal/usecase/upload"
)

func main() {
	// Создаем usecase для обработки загрузки
	deliveryUsecase := delivery_usecase.New("storage/mods.zip")
	listUsecase := list_usecase.New()
	uploadUsecase := upload_usecase.New()

	// Создаем API-сервер с хэндлерами
	srv := server.API{
		Delivery: delivery.NewDownloadHandler(deliveryUsecase),
		Health:   health.New(),
		List:     list.New(listUsecase),
		Upload:   upload.New(uploadUsecase),
	}

	// Создаем сервер Ogen
	ogenServer, err := api.NewServer(srv)
	if err != nil {
		os.Exit(1)
	}

	// Создаем новый Gin роутер с логером
	router := gin.Default()

	// Оборачиваем Ogen в gin.HandlerFunc
	router.NoRoute(gin.WrapH(ogenServer))

	// Запускаем HTTP сервер
	srvAddr := fmt.Sprintf(":%s", "8080")
	httpServer := &http.Server{
		Addr:    srvAddr,
		Handler: router, // Используем Gin как HTTP хэндлер
	}

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		os.Exit(1)
	}
}
