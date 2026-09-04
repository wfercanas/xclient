package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/selsa-inube/xclient-service/config"
	"github.com/selsa-inube/xclient-service/contracts/dto"
	"github.com/selsa-inube/xclient-service/internal/model"
	"github.com/selsa-inube/xclient-service/shared/types"
)

func validateNewCustomer(newCustomer dto.NewCustomer) error {
	if newCustomer.TenandId == nil {
		return errors.New("missing property: tenant_id")
	}

	if newCustomer.CustomerType == "" {
		return errors.New("missing property: customer_type")
	}

	if newCustomer.CustomerType != "person" && newCustomer.CustomerType != "company" {
		return errors.New("invalid property: customer_type should be equal to 'person' or 'company'")
	}

	if newCustomer.CustomerId == "" {
		return errors.New("missing property: customer_id")
	}

	if newCustomer.Name == "" {
		return errors.New("missing property: name")
	}

	if newCustomer.Status == "" {
		return errors.New("missing property: status")
	}

	if newCustomer.Status != "active" && newCustomer.Status != "retired" && newCustomer.Status != "retiring" && newCustomer.Status != "inactive" {
		return errors.New("invalid property: status should be equal to 'active', 'retired', 'retiring' or 'inactive'")
	}

	if newCustomer.AssociationDate == "" {
		return errors.New("missing property: association_date")
	}

	if newCustomer.AssociationDate > types.NewDate(time.Now()) {
		return errors.New("invalid property: association_date should be lower or equal to the present day")
	}

	return nil
}

func CreateNewCustomer(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCustomer dto.NewCustomer

		err := json.NewDecoder(r.Body).Decode(&newCustomer)
		if err != nil {
			app.Error(w, r, http.StatusBadRequest, err)
			return
		}

		err = validateNewCustomer(newCustomer)
		if err != nil {
			app.Error(w, r, http.StatusBadRequest, err)
			return
		}

		customer, err := app.Customer.Create(model.NewCustomer{
			TenantId:        *newCustomer.TenandId,
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
