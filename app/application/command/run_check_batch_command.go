package command

import "github.com/brunobotter/site-sentinel/application/domain"

type RunCheckBatchCommand struct {
	Targets []domain.MonitorTarget
}
