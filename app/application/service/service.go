package service

import (
	"time"

	"github.com/brunobotter/site-sentinel/application/domain"
)

type MonitorPlannerService interface {
	PlanBatch(targets []domain.MonitorTarget) [][]domain.MonitorTarget
}
type CheckAggregationService interface {
	CalculateUptime(results []domain.CheckResult) float64

	AverageLatency(results []domain.CheckResult) float64

	GroupByWindow(
		results []domain.CheckResult,
		window time.Duration,
	) map[time.Time][]domain.CheckResult
}
