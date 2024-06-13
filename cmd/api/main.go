package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/syntaqx/env"
	"go.uber.org/zap"

	"github.com/syntaqx/api/internal/config"
	"github.com/syntaqx/api/internal/handler"
	"github.com/syntaqx/api/internal/router"
	"github.com/syntaqx/api/internal/service"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Load dotenv environment variables
	_ = env.Load()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("unable to initialize logger: %s", err))
	}
	defer logger.Sync()

	cfg := config.NewConfig()

	// Initialize services
	weatherService := service.NewWeatherService(cfg)

	// Initialize handlers
	rootHandler := handler.NewRootHandler()
	healthHandler := handler.NewHealthHandler()
	timeHandler := handler.NewTimeHandler()
	weatherHandler := handler.NewWeatherHandler(weatherService)

	// Initialize router
	r := router.NewRouter(logger)

	// Register routes
	rootHandler.RegisterRoutes(r)
	healthHandler.RegisterRoutes(r)
	timeHandler.RegisterRoutes(r)
	weatherHandler.RegisterRoutes(r)

	// Static files
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	handler.FileServer(r, "/", filesDir)

	srv := &http.Server{
		Addr:    net.JoinHostPort("", cfg.Port),
		Handler: r,
	}

	logger.Info("http server started", zap.String("addr", srv.Addr))
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Error("http server closed unexpectedly", zap.Error(err))
		}
	}

	logger.Info("http server shutting down...")
	srv.Shutdown(ctx)
}
