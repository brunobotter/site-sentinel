package interfaces

import "github.com/seuuser/go-site-monitor/internal/domain"

type CheckRepository interface {
	Save(check domain.Check) error
}
