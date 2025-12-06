-- workflows table
CREATE TABLE IF NOT EXISTS workflows (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    name TEXT NOT NULL,
    trigger_type TEXT NOT NULL,
    webhook_secret TEXT,
    cron_expr TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- workflow_steps table
CREATE TABLE IF NOT EXISTS workflow_steps (
    id SERIAL PRIMARY KEY,
    workflow_id INT NOT NULL REFERENCES workflows(id),
    step_number INT NOT NULL,
    action_type TEXT NOT NULL,
    action_config JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- trigger_events table
CREATE TABLE IF NOT EXISTS trigger_events (
    id SERIAL PRIMARY KEY,
    workflow_id INT NOT NULL REFERENCES workflows(id),
    payload JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- workflow_runs table
CREATE TABLE IF NOT EXISTS workflow_runs (
    id SERIAL PRIMARY KEY,
    workflow_id INT NOT NULL REFERENCES workflows(id),
    trigger_event_id INT NOT NULL REFERENCES trigger_events(id),
    status TEXT NOT NULL,
    error_message TEXT,
    started_at TIMESTAMP,
    finished_at TIMESTAMP
);

-- step_runs table
CREATE TABLE IF NOT EXISTS step_runs (
    id SERIAL PRIMARY KEY,
    workflow_run_id INT NOT NULL REFERENCES workflow_runs(id),
    step_number INT NOT NULL,
    status TEXT NOT NULL,
    input_data JSONB,
    output_data JSONB,
    error_message TEXT,
    started_at TIMESTAMP,
    finished_at TIMESTAMP
);
