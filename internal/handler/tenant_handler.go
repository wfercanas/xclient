package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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
			fmt.Println(err)
			http.Error(w, "failed to create tenant", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "appliation/json")
		json.NewEncoder(w).Encode(dto.Tenant(tenant))
	}
}
