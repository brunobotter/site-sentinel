package repository

import (
	"errors"
	"sync"

	"github.com/seuuser/go-site-monitor/internal/domain"
)

type MemorySiteRepository struct {
	mu    sync.RWMutex
	sites []domain.Site
}

func NewMemorySiteRepository(seed []domain.Site) *MemorySiteRepository {
	copySeed := make([]domain.Site, len(seed))
	copy(copySeed, seed)

	return &MemorySiteRepository{sites: copySeed}
}

func (r *MemorySiteRepository) List() ([]domain.Site, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]domain.Site, len(r.sites))
	copy(result, r.sites)

	return result, nil
}

func (r *MemorySiteRepository) Create(site domain.Site) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, s := range r.sites {
		if s.ID == site.ID {
			return errors.New("site with same id already exists")
		}
	}

	r.sites = append(r.sites, site)
	return nil
}

func (r *MemorySiteRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, s := range r.sites {
		if s.ID == id {
			r.sites = append(r.sites[:i], r.sites[i+1:]...)
			return nil
		}
	}

	return errors.New("site not found")
}
