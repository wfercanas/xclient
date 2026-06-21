package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/selsa-inube/iclient-query-service/contracts/dto"
)

func TestTenantEncodeAndDecode(t *testing.T) {
	initial := &dto.Tenant{
		Id:        1,
		Name:      "acme",
		CreatedAt: time.Now(),
	}

	data, err := json.Marshal(initial)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var final dto.Tenant
	err = json.Unmarshal(data, &final)
	if err != nil {
		t.Fatalf("failed to unmarshall: %v", err)
	}

	if initial.Id != final.Id {
		t.Errorf("expected id %v, got id %v", initial.Id, final.Id)
	}

	if initial.Name != final.Name {
		t.Errorf("expected name %v, got name %v", initial.Name, final.Name)
	}

	if !(initial.CreatedAt.Equal(final.CreatedAt)) {
		t.Errorf("expected created_at %v, got created_at %v", initial.CreatedAt, final.CreatedAt)
	}
}

func TestNewTenantEncodeAndDecode(t *testing.T) {
	initial := dto.NewTenant{
		Name: "acme",
	}

	data, err := json.Marshal(initial)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var final dto.NewTenant
	err = json.Unmarshal(data, &final)
	if err != nil {
		t.Fatalf("failed to unmarshall: %v", err)
	}

	if initial.Name != final.Name {
		t.Errorf("expected %v, got %v", initial.Name, final.Name)
	}
}
