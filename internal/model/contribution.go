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
	return Contribution{
		ProductNumber:      "201-1098717173",
		ProductDescription: "AHORRO PERMANENTE",
		BalanceSavings:     323626,
	}, nil
}
