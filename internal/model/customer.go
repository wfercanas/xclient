package model

import (
	"database/sql"
	"errors"
	"time"

	"github.com/selsa-inube/xclient-service/shared/types"
)

type Customer struct {
	Id              int
	TenantId        int
	CustomerType    string
	CustomerId      string
	Name            string
	Status          string
	AssociationDate types.Date
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type NewCustomer struct {
	TenantId        int
	CustomerType    string
	CustomerId      string
	Name            string
	Status          string
	AssociationDate types.Date
}

type CustomerModel struct {
	DB *sql.DB
}

func (m *CustomerModel) Create(newCustomer NewCustomer) (Customer, error) {
	result := m.DB.QueryRow(`
		INSERT INTO customers (tenant_id, customer_type, customer_id, name, status, association_date)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, tenant_id, customer_type, customer_id, name, status, association_date, created_at, updated_at
	`, newCustomer.TenantId, newCustomer.CustomerType, newCustomer.CustomerId, newCustomer.Name, newCustomer.Status, newCustomer.AssociationDate)

	var customer Customer
	var associationDate time.Time
	err := result.Scan(&customer.Id, &customer.TenantId, &customer.CustomerType, &customer.CustomerId, &customer.Name, &customer.Status, &associationDate, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return Customer{}, err
	}

	customer.AssociationDate = types.NewDate(associationDate)
	return customer, nil
}

func (m *CustomerModel) GetById(customerId int) (Customer, error) {
	result := m.DB.QueryRow(`
		SELECT id, tenant_id, customer_type, customer_id, name, status, association_date, created_at, updated_at
		FROM customers
		WHERE id = $1
	`, customerId)

	var customer Customer
	var associationDate time.Time

	err := result.Scan(&customer.Id, &customer.TenantId, &customer.CustomerType, &customer.CustomerId, &customer.Name, &customer.Status, &associationDate, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Customer{}, sql.ErrNoRows
		} else {
			return Customer{}, err
		}
	}

	customer.AssociationDate = types.NewDate(associationDate)
	return customer, nil
}

func (m *CustomerModel) GetAll(id int, tenantId int, customerType string, customerId string, name string, status string) ([]Customer, error) {
	// results, err := m.DB.Query(`
	// 	SELECT id, tenant_id, customer_type, customer_id, name, status, created_at, updated_at
	// 	FROM customers
	// 	WHERE ($1::int = -1 OR id = $1::int)
	// 	AND ($2::int = -1 OR tenant_id = $2::int)
	// 	AND ($3::text IS NULL OR customer_type ILIKE $3::text)
	// 	AND ($4::text IS NULL OR customer_id ILIKE $4::text)
	// 	AND ($5::text IS NULL OR name ILIKE '%' || $5::text || '%')
	// 	AND ($6::text IS NULL OR status ILIKE $6::text)
	// `, id, tenantId, customerType, customerId, name, status)

	results, err := m.DB.Query(`
		SELECT id, tenant_id, customer_type, customer_id, name, status, association_date, created_at, updated_at
		FROM customers
		WHERE ($1::int = -1 OR id = $1::int)
	`, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	var customers []Customer
	for results.Next() {
		var customer Customer
		var associationDate time.Time
		err := results.Scan(&customer.Id, &customer.TenantId, &customer.CustomerType, &customer.CustomerId, &customer.Name, &customer.Status, &associationDate, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		customer.AssociationDate = types.NewDate(associationDate)
		customers = append(customers, customer)
	}

	return customers, nil
}

func (m *CustomerModel) Delete(customerId int) error {
	_, err := m.DB.Exec(`
		DELETE FROM customers
		WHERE id = $1
	`, customerId)

	return err
}
