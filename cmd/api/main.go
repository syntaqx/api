package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/syntaqx/zapchi"
	"go.uber.org/zap"

	"github.com/syntaqx/api/internal/handler"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("unable to initialize logger: %s", err))
	}
	defer logger.Sync()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	// Initialize handlers
	rootHandler := handler.NewRootHandler()
	timeHandler := handler.NewTimeHandler()

	r := chi.NewRouter()

	// Base middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(zapchi.Logger(logger, "router"))
	r.Use(middleware.Recoverer)

	// Register routes
	rootHandler.RegisterRoutes(r)
	timeHandler.RegisterRoutes(r)

	// Static files
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	handler.FileServer(r, "/", filesDir)

	srv := &http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: r,
	}

	logger.Info("http server started", zap.String("addr", srv.Addr))
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			logger.Error("http server closed unexpectedly", zap.Error(err))
		}
	}
}
