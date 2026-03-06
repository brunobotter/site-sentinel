package domain

import "time"

type Check struct {
	ID        string    `json:"id"`
	SiteID    string    `json:"site_id"`
	Status    string    `json:"status"`
	LatencyMS int64     `json:"latency_ms"`
	CheckedAt time.Time `json:"checked_at"`
}
