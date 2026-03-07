package service

import (
	"time"

	"github.com/brunobotter/site-sentinel/application/domain"
)

type checkAggregationService struct{}

func NewCheckAggregationService() *checkAggregationService {
	return &checkAggregationService{}
}

func (s *checkAggregationService) CalculateUptime(results []domain.CheckResult) float64 {
	if len(results) == 0 {
		return 0
	}

	upCount := 0
	for _, result := range results {
		if result.IsUp {
			upCount++
		}
	}

	return float64(upCount) / float64(len(results)) * 100
}

func (s *checkAggregationService) AverageLatency(results []domain.CheckResult) float64 {
	if len(results) == 0 {
		return 0
	}

	var total time.Duration
	for _, result := range results {
		total += result.ResponseTime
	}

	return total.Seconds() * 1000 / float64(len(results))
}

func (s *checkAggregationService) GroupByWindow(
	results []domain.CheckResult,
	window time.Duration,
) map[time.Time][]domain.CheckResult {
	grouped := make(map[time.Time][]domain.CheckResult)
	if len(results) == 0 {
		return grouped
	}

	if window <= 0 {
		window = time.Minute
	}

	for _, result := range results {
		bucket := result.CheckedAt.Truncate(window)
		grouped[bucket] = append(grouped[bucket], result)
	}

	return grouped
}
