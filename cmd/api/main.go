package main

import (
	"net"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/syntaqx/api/pkg/logger"
)

func main() {
	logger, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	r := http.NewServeMux()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	srv := &http.Server{
		Addr:         net.JoinHostPort("", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Info("http server started", zap.String("port", port))
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			logger.Error("http server closed unexpectedly", zap.Error(err))
		}
	}
}
