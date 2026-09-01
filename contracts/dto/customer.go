package dto

import (
	"time"

	"github.com/selsa-inube/xclient-service/shared/types"
)

type Customer struct {
	Id              int        `json:"id"`
	TenantId        int        `json:"tenant_id"`
	CustomerType    string     `json:"customer_type"`
	CustomerId      string     `json:"customer_id"`
	Name            string     `json:"name"`
	Status          string     `json:"status"`
	AssociationDate types.Date `json:"association_date"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
