package repository

import (
	"fmt"

	"github.com/seuuser/go-site-monitor/internal/domain"
)

type MemoryCheckRepository struct{}

func NewMemoryCheckRepository() *MemoryCheckRepository {
	return &MemoryCheckRepository{}
}

func (r *MemoryCheckRepository) Save(check domain.Check) error {
	fmt.Printf("saved check: site=%s status=%s latency_ms=%d\n", check.SiteID, check.Status, check.LatencyMS)
	return nil
}
