package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshsoftware/peerly-backend/internal/app"
	"github.com/joshsoftware/peerly-backend/internal/pkg/config"
	"github.com/joshsoftware/peerly-backend/internal/pkg/middleware"
)

const (
	versionHeader = "Accept-Version"
	authHeader    = "X-Auth-Token"
)

// NewRouter initializes and returns a new router with the specified dependencies.
func NewRouter(deps app.Dependencies) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	// Version 1 API management
	v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())

	//corevalues
	router.Handle("/core_values/{id:[0-9]+}", middleware.JwtAuthMiddleware(getCoreValueHandler(deps.CoreValueService))).Methods(http.MethodGet).Headers(versionHeader, v1)

	router.Handle("/core_values", middleware.JwtAuthMiddleware(listCoreValuesHandler(deps.CoreValueService))).Methods(http.MethodGet).Headers(versionHeader, v1)

	router.Handle("/core_values", middleware.JwtAuthMiddleware(createCoreValueHandler(deps.CoreValueService))).Methods(http.MethodPost).Headers(versionHeader, v1)

	router.Handle("/core_values/{id:[0-9]+}", middleware.JwtAuthMiddleware(updateCoreValueHandler(deps.CoreValueService))).Methods(http.MethodPut).Headers(versionHeader, v1)
	
	router.Handle("/appreciation", middleware.JwtAuthMiddleware(createAppreciationHandler(deps.AppreciationService))).Methods(http.MethodPost).Headers(versionHeader, v1)

	router.Handle("/appreciation/{id:[0-9]+}", middleware.JwtAuthMiddleware(getAppreciationByIdHandler(deps.AppreciationService))).Methods(http.MethodGet).Headers(versionHeader, v1)

	router.Handle("/appreciations", middleware.JwtAuthMiddleware(getAppreciationsHandler(deps.AppreciationService))).Methods(http.MethodGet).Headers(versionHeader, v1)

	router.Handle("/appreciation/{id:[0-9]+}", middleware.JwtAuthMiddleware(validateAppreciationHandler(deps.AppreciationService))).Methods(http.MethodDelete).Headers(versionHeader, v1)

	return router
}
