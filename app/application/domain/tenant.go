package domain

import "time"

type TenantDomain struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type TenantPage struct {
	Items      []*TenantDomain
	Page       int
	Limit      int
	Total      int
	TotalPages int
}
