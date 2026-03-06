package response

import "time"

type CreateTenantResponse struct {
	TenantId string
	Name     string
	CreateAt time.Time
}
