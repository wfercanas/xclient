package config

import (
	"database/sql"
	"log/slog"

	"github.com/selsa-inube/iclient-query-service/internal/model"
)

type Application struct {
	Logger       *slog.Logger
	Contribution *model.ContributionModel
	Tenant       *model.TenantModel
}

func NewApplication(logger slog.Logger, db *sql.DB) *Application {
	return &Application{
		Logger:       &logger,
		Contribution: &model.ContributionModel{DB: db},
		Tenant:       &model.TenantModel{DB: db},
	}
}
