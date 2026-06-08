package tests

import (
	"encoding/json"
	"testing"

	"github.com/selsa-inube/iclient-query-service/contracts/dto"
)

func TestContributionJSON(t *testing.T) {
	initial := dto.Contribution{
		ContributionName: "APORTES SOCIALES",
		ContributionId:   "123-4567890",
		Balance:          100,
	}

	data, err := json.Marshal(initial)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var final dto.Contribution
	err = json.Unmarshal(data, &final)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if final != initial {
		t.Errorf("expected %+v, got %+v", initial, final)
	}
}
