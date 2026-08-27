package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/selsa-inube/xclient-service/config"
	"github.com/selsa-inube/xclient-service/contracts/dto"
	"github.com/selsa-inube/xclient-service/internal/model"
)

func CreateNewTenant(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newTenant dto.NewTenant

		err := json.NewDecoder(r.Body).Decode(&newTenant)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		tenant, err := app.Tenant.Create(model.NewTenant(newTenant))
		if err != nil {
			app.InternalServerError(w, r, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dto.Tenant(tenant))
	}
}

func GetTenants(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		idStr := params.Get("id")
		name := params.Get("name")

		id := -1
		if idStr != "" {
			var err error
			id, err = strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
		}

		tenants, err := app.Tenant.GetAll(id, name)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			} else {
				app.InternalServerError(w, r, err)
			}
			return
		}

		tenantsDTO := make([]dto.Tenant, len(tenants))
		for i, tenant := range tenants {
			tenantsDTO[i] = dto.Tenant(tenant)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tenantsDTO)
	}
}

func GetTenantById(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tenantId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "invalid tenant id", http.StatusBadRequest)
			return
		}

		tenant, err := app.Tenant.GetById(tenantId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			} else {
				app.InternalServerError(w, r, err)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dto.Tenant(tenant))
	}
}

func DeleteTenant(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tenantId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "invalid tenant id", http.StatusBadRequest)
			return
		}

		err = app.Tenant.Delete(tenantId)
		if err != nil {
			app.InternalServerError(w, r, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
	}
}
