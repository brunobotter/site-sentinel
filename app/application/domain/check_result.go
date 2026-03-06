package domain

import (
	"time"

	"github.com/google/uuid"
)

type CheckResult struct {
	ID           uuid.UUID
	TargetID     uuid.UUID
	StatusCode   int
	ResponseTime time.Duration
	IsUp         bool
	Error        string
	CheckedAt    time.Time
}

func NewCheckResult(
	targetID uuid.UUID,
	statusCode int,
	responseTime time.Duration,
	isUp bool,
	err string,
) *CheckResult {

	return &CheckResult{
		ID:           uuid.New(),
		TargetID:     targetID,
		StatusCode:   statusCode,
		ResponseTime: responseTime,
		IsUp:         isUp,
		Error:        err,
		CheckedAt:    time.Now(),
	}
}
