package handler

import (
	"encoding/json"
	"net/http"

	"github.com/selsa-inube/xclient-service/config"
	"github.com/selsa-inube/xclient-service/contracts/dto"
)

func GetContributionById(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contribution, err := app.Contribution.GetContributionById(r.PathValue("id"))
		if err != nil {
			app.Logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		dtoBeneficiaries := make([]dto.ContributionBeneficiary, len(contribution.Beneficiaries))
		for i, b := range contribution.Beneficiaries {
			dtoBeneficiaries[i] = dto.ContributionBeneficiary(b)
		}

		response := &dto.Contribution{
			ContributionName:          contribution.ProductDescription,
			ContributionId:            contribution.ProductNumber,
			Balance:                   contribution.BalanceSavings,
			ContributionBeneficiaries: dtoBeneficiaries,
		}

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			app.Logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
