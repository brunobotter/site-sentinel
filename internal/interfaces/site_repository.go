package interfaces

import "github.com/seuuser/go-site-monitor/internal/domain"

type SiteRepository interface {
	List() ([]domain.Site, error)
	Create(site domain.Site) error
	Delete(id string) error
}
