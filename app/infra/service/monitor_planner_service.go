package service

import "github.com/brunobotter/site-sentinel/application/domain"

const defaultBatchSize = 10

type monitorPlannerService struct {
	batchSize int
}

func NewMonitorPlannerService(batchSize int) *monitorPlannerService {
	if batchSize <= 0 {
		batchSize = defaultBatchSize
	}

	return &monitorPlannerService{batchSize: batchSize}
}
func (s *monitorPlannerService) PlanBatch(targets []domain.MonitorTarget) [][]domain.MonitorTarget {
	if len(targets) == 0 {
		return [][]domain.MonitorTarget{}
	}

	batches := make([][]domain.MonitorTarget, 0, (len(targets)+s.batchSize-1)/s.batchSize)

	for start := 0; start < len(targets); start += s.batchSize {
		end := start + s.batchSize
		if end > len(targets) {
			end = len(targets)
		}

		batch := make([]domain.MonitorTarget, end-start)
		copy(batch, targets[start:end])
		batches = append(batches, batch)
	}

	return batches
}
