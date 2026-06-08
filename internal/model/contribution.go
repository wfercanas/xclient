package model

import "database/sql"

type Contribution struct {
	ProductNumber      string
	ProductDescription string
	BalanceSavings     float64
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

	return contribution, nil
}
