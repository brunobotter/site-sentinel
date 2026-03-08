package service

import (
	"time"

	"github.com/brunobotter/site-sentinel/application/domain"
)

type checkAggregationService struct{}

// NewCheckAggregationService cria utilitários de agregação para análise de resultados.
func NewCheckAggregationService() *checkAggregationService {
	return &checkAggregationService{}
}

// CalculateUptime calcula o percentual de disponibilidade com base no histórico recebido.
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

// AverageLatency calcula latência média em milissegundos.
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

// GroupByWindow organiza resultados em janelas de tempo para facilitar gráficos/relatórios.
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
