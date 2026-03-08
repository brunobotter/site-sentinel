package controllers

import (
	"errors"
	"strconv"
	"time"

	apihttp "github.com/brunobotter/site-sentinel/api/http"
	"github.com/brunobotter/site-sentinel/api/requests"
	"github.com/brunobotter/site-sentinel/api/response"
	"github.com/brunobotter/site-sentinel/application"
	"github.com/brunobotter/site-sentinel/application/command"
	"github.com/brunobotter/site-sentinel/application/domain"
	"github.com/brunobotter/site-sentinel/application/usecase"
)

type MonitorHandler struct {
	createTargetUseCase     usecase.CreateTargetUseCase
	listTargetsUseCase      usecase.ListTargetsUseCase
	runBatchCheckUseCase    usecase.RunBatchCheckUseCase
	listLatestResultUseCase usecase.ListLatestResultsUseCase
}

// NewMonitorHandler injeta os use cases necessários para os endpoints de monitoria.
//
// Para júnior: controller deve orquestrar entrada/saída HTTP, não conter regra de negócio pesada.
func NewMonitorHandler(
	createTargetUseCase usecase.CreateTargetUseCase,
	listTargetsUseCase usecase.ListTargetsUseCase,
	runBatchCheckUseCase usecase.RunBatchCheckUseCase,
	listLatestResultUseCase usecase.ListLatestResultsUseCase,
) *MonitorHandler {
	return &MonitorHandler{
		createTargetUseCase:     createTargetUseCase,
		listTargetsUseCase:      listTargetsUseCase,
		runBatchCheckUseCase:    runBatchCheckUseCase,
		listLatestResultUseCase: listLatestResultUseCase,
	}
}

// CreateTarget converte payload HTTP em comando da aplicação e delega o cadastro do monitor.
func (h *MonitorHandler) CreateTarget(req *requests.CreateMonitorTargetRequest) *apihttp.HttpResponse {
	cmd := command.CreateTargetCommand{
		URL:            req.URL,
		Name:           req.Name,
		ExpectedStatus: req.ExpectedStatus,
		Timeout:        time.Duration(req.TimeoutMs) * time.Millisecond,
		Interval:       time.Duration(req.IntervalMs) * time.Millisecond,
		Retries:        req.Retries,
		IsActive:       req.IsActive,
	}

	if err := h.createTargetUseCase.Execute(req.Context(), cmd); err != nil {
		return mapApplicationError(err)
	}

	return apihttp.Created(map[string]string{"message": "monitor target criado"})
}

// ListTargets lista os alvos cadastrados e converte o domínio para response DTO.
func (h *MonitorHandler) ListTargets(req apihttp.HttpRequest) *apihttp.HttpResponse {
	targets, err := h.listTargetsUseCase.Execute(req.Context())
	if err != nil {
		return mapApplicationError(err)
	}

	data := make([]response.MonitorTargetResponse, 0, len(targets))
	for _, target := range targets {
		data = append(data, toMonitorTargetResponse(target))
	}

	return apihttp.Ok(data)
}

// RunBatchCheck dispara uma execução manual dos checks para os alvos existentes.
func (h *MonitorHandler) RunBatchCheck(req *requests.RunBatchCheckRequest) *apihttp.HttpResponse {
	targets, err := h.listTargetsUseCase.Execute(req.Context())
	if err != nil {
		return mapApplicationError(err)
	}

	cmd := command.RunCheckBatchCommand{Targets: targets}
	if err := h.runBatchCheckUseCase.Execute(req.Context(), cmd); err != nil {
		return mapApplicationError(err)
	}

	return apihttp.Ok(map[string]string{"message": "batch de checks executado"})
}

// ListLatestResults retorna os checks mais recentes respeitando o limite informado na query string.
func (h *MonitorHandler) ListLatestResults(req *requests.ListMonitorResultsRequest) *apihttp.HttpResponse {
	limit := 50
	if rawLimit := req.QueryParam("limit"); rawLimit != "" {
		parsedLimit, err := strconv.Atoi(rawLimit)
		if err != nil {
			return apihttp.BadRequest("limit invalido")
		}
		limit = parsedLimit
	}

	results, err := h.listLatestResultUseCase.Execute(req.Context(), limit)
	if err != nil {
		return mapApplicationError(err)
	}

	data := make([]response.CheckResultResponse, 0, len(results))
	for _, result := range results {
		data = append(data, toCheckResultResponse(result))
	}

	return apihttp.Ok(data)
}

// toMonitorTargetResponse faz o mapeamento da entidade para o formato público da API.
func toMonitorTargetResponse(target domain.MonitorTarget) response.MonitorTargetResponse {
	return response.MonitorTargetResponse{
		ID:             target.ID.String(),
		Name:           target.Name,
		URL:            target.URL,
		Method:         target.Method,
		ExpectedStatus: target.Policy.ExpectedStatus,
		TimeoutMs:      target.Policy.Timeout.Milliseconds(),
		Retries:        target.Policy.Retries,
		RetryDelayMs:   target.Policy.RetryDelay.Milliseconds(),
		Active:         target.Active,
		CreatedAt:      response.FormatTime(target.CreatedAt),
		UpdatedAt:      response.FormatTime(target.UpdatedAt),
	}
}

// toCheckResultResponse converte o resultado de domínio para o contrato de saída HTTP.
func toCheckResultResponse(result domain.CheckResult) response.CheckResultResponse {
	return response.CheckResultResponse{
		ID:             result.ID.String(),
		TargetID:       result.TargetID.String(),
		StatusCode:     result.StatusCode,
		ResponseTimeMs: result.ResponseTime.Milliseconds(),
		IsUp:           result.IsUp,
		Error:          result.Error,
		CheckedAt:      response.FormatTime(result.CheckedAt),
	}
}

// mapApplicationError traduz erros de negócio para códigos HTTP previsíveis.
//
// Para júnior: centralizar esse mapeamento evita repetição e garante consistência da API.
func mapApplicationError(err error) *apihttp.HttpResponse {
	var validationErr application.ValidationApplicationError
	if errors.As(err, &validationErr) {
		return apihttp.UnprocessableEntity(err.Error())
	}

	var notFoundErr application.NotFoundApplicationError
	if errors.As(err, &notFoundErr) {
		return apihttp.NotFound(err.Error())
	}

	return apihttp.InternalServerError(err.Error())
}

// Health expõe healthcheck específico do módulo de monitoria.
func (h *MonitorHandler) Health(req apihttp.HttpRequest) *apihttp.HttpResponse {
	return apihttp.Ok(map[string]string{"status": "ok", "service": "monitor"})
}
