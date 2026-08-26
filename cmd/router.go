package main

import (
	"net/http"

	"github.com/selsa-inube/iclient-query-service/config"
	"github.com/selsa-inube/iclient-query-service/internal/handler"
)

func router(app *config.Application) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /contributions/{id}", handler.GetContributionById(app))

	mux.HandleFunc("POST /tenants", handler.CreateNewTenant(app))
	mux.HandleFunc("GET /tenants/{$}", handler.GetTenants(app))
	mux.HandleFunc("GET /tenants/{id}", handler.GetTenantById(app))
	mux.HandleFunc("DELETE /tenants/{id}", handler.DeleteTenant(app))

	return mux
}
