BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- targets
CREATE TABLE monitor_targets (
    id UUID PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    url TEXT NOT NULL,
    method VARCHAR(10) NOT NULL DEFAULT 'GET',

    timeout_ms INTEGER NOT NULL DEFAULT 3000,
    expected_status INTEGER NOT NULL DEFAULT 200,
    retries INTEGER NOT NULL DEFAULT 1,
    retry_delay_ms INTEGER NOT NULL DEFAULT 500,

    active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_monitor_targets_active
ON monitor_targets(active);

-- results
CREATE TABLE check_results (
    id UUID PRIMARY KEY,

    target_id UUID NOT NULL,

    status_code INTEGER,
    response_time_ms INTEGER,
    is_up BOOLEAN,

    error TEXT,

    checked_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_check_results_target
        FOREIGN KEY (target_id)
        REFERENCES monitor_targets(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_check_results_target_checked
ON check_results(target_id, checked_at DESC);

CREATE INDEX idx_check_results_checked
ON check_results(checked_at DESC);

COMMIT;