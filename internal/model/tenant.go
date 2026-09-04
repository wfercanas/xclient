package model

import (
	"database/sql"
	"errors"
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

func (m *TenantModel) GetAll(id int, name string) ([]Tenant, error) {
	results, err := m.DB.Query(`
		SELECT id, name, created_at
		FROM tenants
		WHERE ($1::int = -1 OR id = $1::int)
		AND ($2::text IS NULL OR name ILIKE '%' || $2::text || '%')
	`, id, name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		} else {
			return nil, err
		}
	}

	var tenants []Tenant
	for results.Next() {
		var tenant Tenant

		err := results.Scan(&tenant.Id, &tenant.Name, &tenant.CreatedAt)
		if err != nil {
			return nil, err
		}

		tenants = append(tenants, tenant)
	}

	return tenants, nil
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

	return err
}
