package router

import (
	"fmt"
	"net"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/syntaqx/api/docs"
	_ "github.com/syntaqx/api/docs"
	"github.com/syntaqx/api/internal/config"
	"github.com/syntaqx/zapchi"
	"go.uber.org/zap"
)

// @title           Syntaqx Personal API
// @version         1.0
// @description     My Personal API

// @license.name  MIT
// @license.url   https://opensource.org/license/mit

// @host      localhost:8080
// @BasePath  /

// NewRouter creates a new chi router with base middleware and swagger docs
func NewRouter(config *config.Config, logger *zap.Logger) chi.Router {
	r := chi.NewRouter()

	// Base middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(zapchi.Logger(logger, "router"))
	r.Use(middleware.Recoverer)

	host := net.JoinHostPort(config.Host, config.Port)

	docs.SwaggerInfo.Host = host

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", host)), //The url pointing to API definition
	))

	return r
}
