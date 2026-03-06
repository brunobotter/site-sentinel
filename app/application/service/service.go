package service

import (
	"time"

	"github.com/brunobotter/site-sentinel/application/domain"
)

// Ele é responsável por planejar como os checks serão executados.
// Como devemos organizar os targets para execução?
// Lista de MonitorTarget e divide em lotes
type MonitorPlannerService interface {
	PlanBatch(targets []domain.MonitorTarget) [][]domain.MonitorTarget
}

// Esse serviço trabalha depois da execução dos checks.
// resultados de monitoramento
// O CheckAggregationService transforma dados brutos em métricas.
type CheckAggregationService interface {
	//Calcular uptime
	CalculateUptime(results []domain.CheckResult) float64

	//calcular latencia media
	AverageLatency(results []domain.CheckResult) float64

	//Agrupar resultados por janela de tempo
	GroupByWindow(
		results []domain.CheckResult,
		window time.Duration,
	) map[time.Time][]domain.CheckResult
}
