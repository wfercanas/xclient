package handler

import (
	"encoding/json"
	"net/http"

	"github.com/selsa-inube/xclient-service/config"
	"github.com/selsa-inube/xclient-service/contracts/dto"
	"github.com/selsa-inube/xclient-service/internal/model"
)

func CreateNewCustomer(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCustomer dto.NewCustomer

		err := json.NewDecoder(r.Body).Decode(&newCustomer)
		if err != nil {
			app.Error(w, r, http.StatusBadRequest, err)
			return
		}

		customer, err := app.Customer.Create(model.NewCustomer{
			TenantId:        newCustomer.TenandId,
			CustomerType:    newCustomer.CustomerType,
			CustomerId:      newCustomer.CustomerId,
			Name:            newCustomer.Name,
			Status:          newCustomer.Status,
			AssociationDate: newCustomer.AssociationDate,
		})
		if err != nil {
			app.Error(w, r, http.StatusInternalServerError, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dto.Customer(customer))
	}
}
