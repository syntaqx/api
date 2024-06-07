package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/syntaqx/api/internal/handler"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	// Initialize handlers
	rootHandler := handler.NewRootHandler()
	timeHandler := handler.NewTimeHandler()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	rootHandler.RegisterRoutes(r)
	timeHandler.RegisterRoutes(r)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	handler.FileServer(r, "/", filesDir)

	srv := &http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: r,
	}

	fmt.Printf("http listenning on port %s\n", port)
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			fmt.Printf("http server closed unexpectedly: %v\n", err)
		}
	}
}
