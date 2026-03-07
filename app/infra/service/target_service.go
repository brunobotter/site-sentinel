package service

import (
	"context"
	"strings"
	"time"

	"github.com/brunobotter/site-sentinel/application/command"
	"github.com/brunobotter/site-sentinel/application/domain"
	"github.com/brunobotter/site-sentinel/application/repo"
	"github.com/brunobotter/site-sentinel/application/validator"
)

type targetService struct {
	targetRepo repo.MonitorTargetRepository
}

func NewTargetService(targetRepo repo.MonitorTargetRepository) *targetService {
	return &targetService{targetRepo: targetRepo}
}

func (s *targetService) Create(ctx context.Context, cmd command.CreateTargetCommand) error {
	fieldValidator := validator.NewFieldValidatorControl()
	fieldValidator.AddFieldValidator("url", cmd.URL, validator.Required())
	fieldValidator.AddFieldValidator("name", cmd.Name, validator.Required())
	fieldValidator.AddFieldValidator("expectedStatus", cmd.ExpectedStatus, validator.MinNumber(100), validator.MaxNumber(599))
	if err := fieldValidator.Error(); err != nil {
		return err
	}

	policy := domain.DefaultCheckPolicy()
	policy.ExpectedStatus = cmd.ExpectedStatus
	policy.Retries = cmd.Retries
	if cmd.Timeout > 0 {
		policy.Timeout = cmd.Timeout
	}
	if cmd.Interval > 0 {
		policy.RetryDelay = cmd.Interval
	}

	target := domain.NewMonitorTarget(strings.TrimSpace(cmd.Name), strings.TrimSpace(cmd.URL), "GET", policy)
	target.Active = cmd.IsActive
	target.UpdatedAt = time.Now()

	return s.targetRepo.Create(ctx, *target)
}

func (s *targetService) List(ctx context.Context) ([]domain.MonitorTarget, error) {
	return s.targetRepo.List(ctx, 100, 0)
}
