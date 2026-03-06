# go-site-monitor (starter profissional)

Estrutura profissional mínima para iniciar o servidor com separação de camadas (Clean Architecture style), sem implementar ainda a rotina de concorrência do monitoramento.

## Estrutura do projeto

```text
cmd/
  server/
    main.go

internal/
  domain/
    site.go
    check.go
  interfaces/
    site_repository.go
    check_repository.go
  repository/
    memory_site_repository.go
    memory_check_repository.go
  service/
    site_service.go
  monitoring/
    README.md
  delivery/
    http/
      router.go
      site_handler.go

pkg/
  logger/
    logger.go
```

## Filosofia da arquitetura

Fluxo principal atual:

```text
delivery (HTTP/API)
      ↓
service (regras de negócio)
      ↓
interfaces (contratos)
      ↓
repository (implementação)
      ↓
domain (entidades)
```

`monitoring` já está reservado para evoluir depois (checker, worker pool, scheduler e result processor), sem acoplamento com HTTP.

## Como rodar

```bash
go mod tidy
go run cmd/server/main.go
```

Servidor sobe em `http://localhost:8080`.

## Endpoints iniciais (MVP de API)

- `GET /health`
- `GET /sites`
- `POST /sites`
- `DELETE /sites/{id}`

Exemplo de criação:

```bash
curl -X POST http://localhost:8080/sites \
  -H 'Content-Type: application/json' \
  -d '{"url":"https://google.com","interval_seconds":60}'
```

## Roadmap por fases

### Fase 1 — Setup do projeto (base)
- Base limpa em camadas.
- Servidor HTTP no ar.
- Repositório inicial em memória.

### Fase 2 — CRUD de sites
- `POST /sites`
- `GET /sites`
- `DELETE /sites/:id`
- Evoluir de memória para PostgreSQL.

### Fase 3 — Primeiro checker
- Implementar `CheckSite(url)` com latência + status.

### Fase 4 — Worker Pool
- Adicionar jobs/results + workers.

### Fase 5 — Scheduler simples
- Tick periódico para buscar sites e enviar jobs.

### Fase 6 — Result processor
- Consumir resultados e persistir checks.

### Fase 7 — Dashboard API
- `GET /sites/status`
- `GET /sites/:id/history`

### Fase 8 — Frontend React
- Vite + Tailwind + Recharts.

### Fase 9 — WebSocket
- Atualização em tempo real para dashboard.

### Fase 10 — Scheduler eficiente
- Uso de `next_check_at` para scheduling incremental.

### Fase 11 — HTTP client otimizado
- Reuso de conexões e timeout ajustado.

### Fase 12 — Sistema distribuído
- Separar API, scheduler, fila e workers.
- Redis/asynq para enfileiramento.

### Fase 13 — Escala 50k+
- Multiplicar nós de workers + tuning de pool.

### Fase 14 — Alertas
- Email, Telegram, Webhook.

### Fase 15 — Métricas
- Prometheus + Grafana.

## Próximo passo natural

Trocar `repository` em memória por PostgreSQL, mantendo as interfaces já definidas para não quebrar `service` e `delivery`.
