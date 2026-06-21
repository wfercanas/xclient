package dto

import "time"

type Tenant struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type NewTenant struct {
	Name string `json:"name"`
}
