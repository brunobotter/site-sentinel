# Site Sentinel

Aplicação em Go para monitorar disponibilidade de URLs e exibir painel web.

## Pré-requisitos

- Go 1.25+
- (Opcional) Docker + Docker Compose para subir banco/redis rapidamente

## Rodando local no VS Code (sem Docker)

> Abrir o VS Code **não inicia o servidor automaticamente**. É necessário executar a aplicação (Run/Debug ou terminal).

1. Entre na pasta do app:
   ```bash
   cd app
   ```
2. Inicie o servidor:
   ```bash
   PROFILE=dev go run ./main.go
   ```
3. Verifique no log a porta usada, por exemplo:
   - `http server started on [::]:8880`
4. Abra no navegador exatamente a mesma porta do log:
   - `http://localhost:8880/app/` (exemplo)

### Erro `ERR_TOO_MANY_REDIRECTS`

Se aparecer esse erro:

- Confira se você está acessando a porta correta (a mesma do log).
- Evite usar outra porta que já esteja ocupada por outro serviço (ex.: `8080`).
- Limpe cookies/site data do `localhost` no navegador e recarregue.

## Rodando com Docker Compose

```bash
docker compose up -d --build
```

Serviço web/API em:

- `http://localhost:8080/app/`

## Healthcheck

```bash
curl -sS http://localhost:8080/health
```
