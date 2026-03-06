package domain

import (
	"time"

	"github.com/google/uuid"
)

type MonitorTarget struct {
	ID        uuid.UUID
	Name      string
	URL       string
	Method    string
	Policy    CheckPolicy
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewMonitorTarget(
	name string,
	url string,
	method string,
	policy CheckPolicy,
) *MonitorTarget {

	now := time.Now()

	return &MonitorTarget{
		ID:        uuid.New(),
		Name:      name,
		URL:       url,
		Method:    method,
		Policy:    policy,
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (t *MonitorTarget) Disable() {
	t.Active = false
	t.UpdatedAt = time.Now()
}

func (t *MonitorTarget) Enable() {
	t.Active = true
	t.UpdatedAt = time.Now()
}
