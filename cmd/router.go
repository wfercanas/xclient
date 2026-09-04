package main

import (
	"net/http"

	"github.com/selsa-inube/xclient-service/config"
	"github.com/selsa-inube/xclient-service/internal/handler"
	"github.com/selsa-inube/xclient-service/middlewares"
)

func router(app *config.Application) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /contributions/{id}", handler.GetContributionById(app))

	mux.HandleFunc("POST /tenants", handler.CreateNewTenant(app))
	mux.HandleFunc("GET /tenants/{$}", handler.GetTenants(app))
	mux.HandleFunc("GET /tenants/{id}", handler.GetTenantById(app))
	mux.HandleFunc("DELETE /tenants/{id}", handler.DeleteTenant(app))

	mux.HandleFunc("POST /customers", handler.CreateNewCustomer(app))
	mux.HandleFunc("GET /customers/{id}", handler.GetCustomerById(app))
	mux.HandleFunc("DELETE /customers/{id}", handler.DeleteCustomer(app))

	return middlewares.LogRequests(app, mux)
}
