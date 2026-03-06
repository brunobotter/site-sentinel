#!/usr/bin/env bash
set -euo pipefail

# ============================
# Configuração
# ============================
: "${DATABASE_URL?DATABASE_URL nao definido}"
export PGCONNECT_TIMEOUT=3

echo "[migrator] aguardando banco..."

# ============================
# Aguarda Postgres ficar pronto
# ============================
ready=0
for i in $(seq 1 60); do
  if psql "$DATABASE_URL" -tAc "select 1" >/dev/null 2>&1; then
    echo "[migrator] banco disponível"
    ready=1
    break
  fi
  sleep 1
done

if [[ "$ready" != "1" ]]; then
  echo "[migrator] erro: banco não ficou disponível a tempo"
  exit 1
fi

# ============================
# Tabela de controle
# ============================
psql "$DATABASE_URL" -v ON_ERROR_STOP=1 <<'SQL'
CREATE TABLE IF NOT EXISTS schema_migrations (
  filename   text PRIMARY KEY,
  applied_at timestamptz NOT NULL DEFAULT now()
);
SQL

# ============================
# Coleta e ordena migrations
# ============================
shopt -s nullglob
files=(/migrations/*.sql)

if [[ ${#files[@]} -eq 0 ]]; then
  echo "[migrator] sem migrations"
  exit 0
fi

# Ordenação correta e segura
mapfile -t files < <(printf "%s\n" "${files[@]}" | sort -V)

# ============================
# Executa migrations
# ============================
for fsql in "${files[@]}"; do
  b="$(basename "$fsql")"

applied="$(
  psql "$DATABASE_URL" -tAc \
    "select 1 from schema_migrations where filename = '$(printf "%s" "$b" | sed "s/'/''/g")' limit 1" \
    | tr -d '[:space:]'
)"

  if [[ "$applied" == "1" ]]; then
    echo "[migrator] skip $b"
    continue
  fi

  echo "[migrator] apply $b"
  psql "$DATABASE_URL" -v ON_ERROR_STOP=1 -f "$fsql"

  psql "$DATABASE_URL" -v ON_ERROR_STOP=1 \
    -c "insert into schema_migrations(filename) values ('$b');"
done

echo "[migrator] ok"
