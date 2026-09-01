package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/selsa-inube/xclient-service/contracts/dto"
	"github.com/selsa-inube/xclient-service/shared/types"
)

func TestCustomerEncodeAndDecode(t *testing.T) {
	initial := &dto.Customer{
		Id:              1,
		TenantId:        2,
		CustomerType:    "person",
		CustomerId:      "1010222333",
		Name:            "John Doe",
		Status:          "active",
		AssociationDate: types.NewDate(time.Now()),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	data, err := json.Marshal(initial)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var final dto.Customer
	err = json.Unmarshal(data, &final)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if initial.Id != final.Id {
		t.Errorf("expected id %v, got id %v", initial.Id, final.Id)
	}

	if initial.TenantId != final.TenantId {
		t.Errorf("expected tenant id %v, got tenant id %v", initial.TenantId, final.TenantId)
	}

	if initial.CustomerType != final.CustomerType {
		t.Errorf("expected customer type %v, got customer type %v", initial.CustomerType, final.CustomerType)
	}

	if initial.CustomerId != final.CustomerId {
		t.Errorf("expected customer id %v, got customer id %v", initial.CustomerId, final.CustomerId)
	}

	if initial.Name != final.Name {
		t.Errorf("expected name %v, got name %v", initial.Name, final.Name)
	}

	if initial.Status != final.Status {
		t.Errorf("expected status %v, got status %v", initial.Status, final.Status)
	}

	if initial.AssociationDate != final.AssociationDate {
		t.Errorf("expected association date %v, got association date %v", initial.AssociationDate, final.AssociationDate)
	}

	if !initial.CreatedAt.Equal(final.CreatedAt) {
		t.Errorf("expected created at %v, got created at %v", initial.CreatedAt, final.CreatedAt)
	}

	if !initial.UpdatedAt.Equal(final.UpdatedAt) {
		t.Errorf("expected updated at %v, got created at %v", initial.UpdatedAt, final.UpdatedAt)
	}
}
