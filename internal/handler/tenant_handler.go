package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/selsa-inube/iclient-query-service/config"
	"github.com/selsa-inube/iclient-query-service/contracts/dto"
	"github.com/selsa-inube/iclient-query-service/internal/model"
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
			http.Error(w, "failed to create tenant", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dto.Tenant(tenant))
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
				http.Error(w, "failed to get tenant", http.StatusInternalServerError)
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
			http.Error(w, "failed to delete tenant", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
	}
}
