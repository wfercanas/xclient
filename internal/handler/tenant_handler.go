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
			app.Error(w, r, http.StatusBadRequest, err)
			return
		}

		tenant, err := app.Tenant.Create(model.NewTenant(newTenant))
		if err != nil {
			app.Error(w, r, http.StatusInternalServerError, err)
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
				app.Error(w, r, http.StatusBadRequest, err)
				return
			}
		}

		tenants, err := app.Tenant.GetAll(id, name)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				app.Error(w, r, http.StatusNotFound, err)
			} else {
				app.Error(w, r, http.StatusInternalServerError, err)
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
			app.Error(w, r, http.StatusBadRequest, err)
			return
		}

		tenant, err := app.Tenant.GetById(tenantId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				app.Error(w, r, http.StatusNotFound, err)
			} else {
				app.Error(w, r, http.StatusInternalServerError, err)
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
			app.Error(w, r, http.StatusBadRequest, err)
			return
		}

		err = app.Tenant.Delete(tenantId)
		if err != nil {
			app.Error(w, r, http.StatusInternalServerError, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
	}
}
