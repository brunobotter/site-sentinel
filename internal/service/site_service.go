package service

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/seuuser/go-site-monitor/internal/domain"
	"github.com/seuuser/go-site-monitor/internal/interfaces"
)

type SiteService struct {
	repo interfaces.SiteRepository
}

func NewSiteService(repo interfaces.SiteRepository) *SiteService {
	return &SiteService{repo: repo}
}

func (s *SiteService) ListSites() ([]domain.Site, error) {
	return s.repo.List()
}

func (s *SiteService) CreateSite(rawURL string, intervalSeconds int) (domain.Site, error) {
	if intervalSeconds <= 0 {
		return domain.Site{}, errors.New("interval_seconds must be greater than zero")
	}

	parsed, err := url.ParseRequestURI(rawURL)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return domain.Site{}, errors.New("invalid URL")
	}

	site := domain.Site{
		ID:              fmt.Sprintf("%d", time.Now().UnixNano()),
		URL:             strings.TrimSpace(rawURL),
		IntervalSeconds: intervalSeconds,
	}

	if err := s.repo.Create(site); err != nil {
		return domain.Site{}, err
	}

	return site, nil
}

func (s *SiteService) DeleteSite(id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("id is required")
	}
	return s.repo.Delete(id)
}
