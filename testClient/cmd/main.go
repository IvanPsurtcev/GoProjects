package main

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"testClient/config"
	"testClient/internal/client/delivery/http"
	"testClient/internal/client/usecase"
	"testClient/internal/service"
	"testClient/pkg/slice"
)

func main() {
	cfg := config.CreateConfig()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	service := &service.MockService{}
	safeSlice := slice.NewSafeSlice()

	clientUC := usecase.NewClient(service, safeSlice)
	clientHandler := http.NewHandlers(logger, clientUC)

	r := router.New()
	clientGroup := r.Group("/client")
	http.InitRouter(clientGroup, clientHandler)
	serverAddr := fmt.Sprintf("%s:%s", cfg.System.Host, cfg.System.Port)
	logger.Fatal("Failed to run server", zap.Error(fasthttp.ListenAndServe(serverAddr, r.Handler)))
}
