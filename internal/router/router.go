package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "github.com/syntaqx/api/docs"
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
func NewRouter(logger *zap.Logger) chi.Router {
	r := chi.NewRouter()

	// Base middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(zapchi.Logger(logger, "router"))
	r.Use(middleware.Recoverer)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	return r
}
