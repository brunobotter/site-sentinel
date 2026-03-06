# Roadmap — Site Sentinel (baseado na estrutura atual)

Este roadmap usa **o modelo de pastas já existente** no projeto:

- `app/api` (controllers, requests, middlewares, viewmodels)
- `app/application` (domain, service, usecase, repo interfaces)
- `app/infra` (implementações externas: http, repo, logger)
- `app/main` (bootstrap, providers, server, DI/container)
- `app/util` (itens compartilhados)

## Objetivo técnico

Escalar o monitoramento para **10k+ alvos** com foco em:

- alta concorrência (goroutines + worker pool)
- uso eficiente de CPU/memória
- controle de taxa (rate limit)
- observabilidade (métricas)
- confiabilidade (timeouts, retries, circuit breaker opcional)

## Fase 1 — Definir domínio de monitoramento

### Onde implementar

- `app/application/domain`
- `app/application/command`

### Entregas

1. Criar entidades:Feito
   - `monitor_target.go`
   - `check_result.go`
   - `check_policy.go` (timeout, expected status, retries)
2. Criar comandos para casos de uso de monitoramento: FEITO
   - `CreateTargetCommand`
   - `RunCheckBatchCommand`

## Fase 2 — Contratos (repositórios e serviços)

### Onde implementar

- `app/application/repo`
- `app/application/service`
- `app/application/usecase`

### Entregas

1. Interfaces de repositório:
   - `MonitorTargetRepository`
   - `CheckResultRepository`
2. Serviços de domínio:
   - `MonitorPlannerService` (planeja lotes de execução)
   - `CheckAggregationService` (estatísticas por janela)
3. Use cases:
   - `CreateTargetUseCase`
   - `ListTargetsUseCase`
   - `RunBatchCheckUseCase`
   - `ListLatestResultsUseCase`

## Fase 3 — Checker HTTP de alta performance

### Onde implementar

- `app/infra/http`
- `app/application/http` (constantes/contratos)

### Entregas

1. Cliente HTTP com:
   - connection pooling ajustado
   - timeout por request
   - keep-alive habilitado
   - limite por host
2. Função de check:
   - mede latência
   - valida status esperado
   - captura erro (DNS, timeout, TLS etc.)

## Fase 4 — Worker pool e pipeline concorrente

### Onde implementar

- `app/application/service` (coordenação)
- `app/main/providers` (injeção/configuração)

### Entregas

1. Estruturar pipeline:
   - `jobs chan MonitorTarget`
   - `results chan CheckResult`
2. Worker pool configurável (ex.: 300–1000 workers)
3. Backpressure com filas limitadas
4. Shutdown gracioso (context cancellation)

## Fase 5 — Scheduler periódico

### Onde implementar

- `app/application/service`
- `app/main/providers`

### Entregas

1. Agendador em intervalo fixo (ex.: 30s/60s)
2. Carrega targets ativos e envia para o pool
3. Evita sobreposição de ciclos (lock de execução)

## Fase 6 — Persistência e tuning de banco

### Onde implementar

- `app/infra/repo`
- `migrator`

### Entregas

1. Tabelas:
   - monitor_targets
   - check_results (particionamento por data opcional)
2. Índices para consultas de dashboard:
   - `(target_id, checked_at desc)`
   - `(checked_at desc)`
3. Batch insert de resultados para reduzir overhead

## Fase 7 — API de monitoramento

### Onde implementar

- `app/api/controllers`
- `app/api/requests`
- `app/api/response`
- `app/main/server/router`

### Entregas

1. Endpoints:
   - `POST /monitor/targets`
   - `GET /monitor/targets`
   - `GET /monitor/results`
   - `GET /monitor/health`
2. Validação de payload e paginação
3. Reuso dos middlewares já existentes

## Fase 8 — Rate limit e proteção

### Onde implementar

- `app/api/middlewares`
- `app/application/service`

### Entregas

1. Rate limit global e por host
2. Limite de concorrência por domínio
3. Estratégia de retry com jitter para falhas transitórias

## Fase 9 — Métricas e observabilidade

### Onde implementar

- `app/infra/logger`
- `app/api/controllers` (endpoint de métricas)

### Entregas

1. Métricas principais:
   - checks_total
   - checks_failed_total
   - check_latency_ms
   - active_workers
   - queue_depth
2. Logs estruturados por execução de check
3. Health e readiness completos

## Fase 10 — Meta 10k+ alvos (benchmark e ajustes)

### Carga alvo inicial

- 10.000 alvos
- janela de 60s
- ~166 checks/s

### Configuração inicial recomendada

- `workers = 500`
- `job_queue = 20_000`
- `timeout = 3s`
- `max_idle_conns` e `max_conns_per_host` ajustados por ambiente

### Critérios de sucesso

- ciclo completo em <= 60s
- erro por timeout controlado
- uso de memória estável sob carga contínua

---

## Próximo passo imediato (prático)

1. Criar domínio e contratos de monitoramento em `app/application`.
2. Implementar checker HTTP em `app/infra/http`.
3. Subir worker pool + scheduler por provider em `app/main/providers`.
4. Expor primeiros endpoints de monitor em `app/api/controllers`.

Com isso você já sai do setup e entra em execução concorrente real no padrão arquitetural que seu projeto já usa.
