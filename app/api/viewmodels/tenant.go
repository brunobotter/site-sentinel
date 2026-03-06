package viewmodels

import (
	"time"

	"github.com/brunobotter/site-sentinel/application/domain"
)

type TenantViewModel struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewTenantViewModel(t *domain.TenantDomain) TenantViewModel {
	return TenantViewModel{
		Id:        t.Id,
		Name:      t.Name,
		CreatedAt: FormatBRDateTime(t.CreatedAt),
		UpdatedAt: FormatBRDateTime(t.UpdatedAt),
	}
}

type TenantPageViewModel struct {
	Data       []TenantViewModel `json:"data"`
	Page       int               `json:"page"`
	Limit      int               `json:"limit"`
	Total      int               `json:"total"`
	TotalPages int               `json:"total_pages"`
}

func NewTenantListViewModel(tenants []*domain.TenantDomain) []TenantViewModel {
	list := make([]TenantViewModel, 0, len(tenants))

	for _, t := range tenants {
		list = append(list, NewTenantViewModel(t))
	}

	return list
}
func NewTenantPageViewModel(tenants *domain.TenantPage) TenantPageViewModel {
	return TenantPageViewModel{
		Data:       NewTenantListViewModel(tenants.Items),
		Page:       tenants.Page,
		Limit:      tenants.Limit,
		Total:      tenants.Total,
		TotalPages: tenants.TotalPages,
	}
}

// dd/MM/yyyy HH:mm:ss
func FormatBRDateTime(dt time.Time) string {
	return dt.Format("02/01/2006 15:04:05")
}
