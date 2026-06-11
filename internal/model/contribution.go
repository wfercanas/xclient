package model

import (
	"database/sql"
)

type Contribution struct {
	ProductNumber      string
	ProductDescription string
	BalanceSavings     float64
	Beneficiaries      []ContributionBeneficiary
}

type ContributionModel struct {
	DB *sql.DB
}

func (m *ContributionModel) GetContributionById(id string) (Contribution, error) {
	result := m.DB.QueryRow(`
		SELECT product_number, product_description, balance_savings
		FROM fondecom.saving_product_catalog
		WHERE product_number = $1`, id)

	var contribution Contribution
	err := result.Scan(&contribution.ProductNumber, &contribution.ProductDescription, &contribution.BalanceSavings)
	if err != nil {
		return Contribution{}, err
	}

	results, err := m.DB.Query(`
		SELECT b.beneficiary_name, b.benefit_percentage
		FROM fondecom.saving_beneficiary b
		JOIN fondecom.saving_product_catalog s
		ON b.product_id = s.product_id
		WHERE s.product_number = $1
	`, id)
	if err != nil {
		return Contribution{}, err
	}
	defer results.Close()

	var beneficiaries []ContributionBeneficiary
	for results.Next() {
		var beneficiary ContributionBeneficiary
		err := results.Scan(&beneficiary.BeneficiaryName, &beneficiary.Percentage)
		if err != nil {
			return Contribution{}, err
		}
		beneficiaries = append(beneficiaries, beneficiary)
	}

	err = results.Err()
	if err != nil {
		return Contribution{}, err
	}

	contribution.Beneficiaries = beneficiaries
	return contribution, nil
}
