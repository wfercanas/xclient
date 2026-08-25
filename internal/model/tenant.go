package model

import (
	"database/sql"
	"time"
)

type NewTenant struct {
	Name string
}

type Tenant struct {
	Id        int
	Name      string
	CreatedAt time.Time
}

type TenantModel struct {
	DB *sql.DB
}

func (m *TenantModel) Create(newTenant NewTenant) (Tenant, error) {
	result := m.DB.QueryRow(`
		INSERT INTO tenants (name)
		VALUES ($1)
		RETURNING id, name, created_at
	`, newTenant.Name)

	var tenant Tenant
	err := result.Scan(&tenant.Id, &tenant.Name, &tenant.CreatedAt)
	if err != nil {
		return Tenant{}, err
	}

	return tenant, nil
}

func (m *TenantModel) GetById(id int) (Tenant, error) {
	result := m.DB.QueryRow(`
		SELECT id, name, created_at
		FROM tenants
		WHERE id = $1
	`, id)

	var tenant Tenant
	err := result.Scan(&tenant.Id, &tenant.Name, &tenant.CreatedAt)
	if err != nil {
		return Tenant{}, err
	}

	return tenant, nil
}

func (m *TenantModel) Delete(id int) error {
	_, err := m.DB.Exec(`
		DELETE FROM tenants
		WHERE id = $1
	`, id)

	if err != nil {
		return err
	}

	return nil
}
