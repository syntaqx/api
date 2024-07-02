package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/syntaqx/env"
	"go.uber.org/zap"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/syntaqx/api/internal/config"
	"github.com/syntaqx/api/internal/handler"
	"github.com/syntaqx/api/internal/model"
	"github.com/syntaqx/api/internal/repository/memory"
	"github.com/syntaqx/api/internal/repository/postgres"
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

	// Initialize config
	cfg := config.NewConfig()

	// Initialize database
	db, err := gorm.Open(pg.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to PostgreSQL database: ", err)
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Initialize repositories
	userRepository := postgres.NewUserRepository(db)
	secretRepository := memory.NewSecretRepository()

	// Initialize services
	weatherService := service.NewWeatherService(cfg)
	userService := service.NewUserService(userRepository)
	secretService := service.NewSecretService(secretRepository)

	// Initialize handlers
	rootHandler := handler.NewRootHandler()
	healthHandler := handler.NewHealthHandler()
	timeHandler := handler.NewTimeHandler()
	weatherHandler := handler.NewWeatherHandler(weatherService)
	usersHandler := handler.NewUsersHandler(userService)
	gamesHandler := handler.NewGamesHandler()
	bookmarksHandler := handler.NewBookmarksHandler()
	secretHandler := handler.NewSecretHandler(secretService)

	// Initialize router
	r := router.NewRouter(cfg, logger)

	// Register routes
	rootHandler.RegisterRoutes(r)
	healthHandler.RegisterRoutes(r)
	timeHandler.RegisterRoutes(r)
	weatherHandler.RegisterRoutes(r)
	usersHandler.RegisterRoutes(r)
	gamesHandler.RegisterRoutes(r)
	bookmarksHandler.RegisterRoutes(r)
	secretHandler.RegisterRoutes(r)

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
