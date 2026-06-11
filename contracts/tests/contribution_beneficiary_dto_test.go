package tests

import (
	"encoding/json"
	"testing"

	"github.com/selsa-inube/iclient-query-service/contracts/dto"
)

func TestContributionBeneficiaryJSON(t *testing.T) {
	initial := dto.ContributionBeneficiary{
		BeneficiaryName: "John Doe",
		Percentage:      "51%",
	}

	data, err := json.Marshal(initial)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var final dto.ContributionBeneficiary
	err = json.Unmarshal(data, &final)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if final != initial {
		t.Errorf("expected %+v, got %+v", initial, final)
	}
}
