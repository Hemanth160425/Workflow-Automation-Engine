Workflow Automation Engine (Zapier-like Backend)

A production-grade workflow automation engine built in Go that executes multi-step workflows asynchronously using Redis queues and PostgreSQL for durable state tracking.

This project demonstrates how real systems like Zapier, n8n, Temporal, GitHub Actions work internally.

🚀 What This Project Does

Executes workflows asynchronously using background workers

Supports multi-step workflows with input → output chaining

Provides durable execution state with per-workflow and per-step logs

Uses Redis as a job queue for scalability and concurrency

Uses PostgreSQL for reliability, retries, and observability

Clean separation between API server and Worker engine

This is not a CRUD app — it is a backend orchestration engine.

🧠 Core Concepts

Workflow – Static definition (trigger + ordered steps)

Trigger Event – Input payload that starts execution

Workflow Run – One execution of a workflow

Step Run – Execution record of each step

Worker – Background process that executes workflows

Queue – Redis-backed job queue for async execution

🏗️ Architecture Overview
API Server (Gin)
   └─ stores workflows, steps, trigger events
   └─ pushes jobs to Redis

Redis Queue
   └─ holds workflow execution jobs

Worker (Go)
   └─ runs forever
   └─ pops jobs from Redis
   └─ executes workflow step-by-step
   └─ writes execution state to PostgreSQL

📦 Tech Stack

Language: Go

HTTP Framework: Gin

Database: PostgreSQL (sqlx)

Queue: Redis

Infra: Docker, Docker Compose

📁 Project Structure
workflow-engine/
├── cmd/
│   ├── api/        # API server (HTTP)
│   └── worker/     # Background worker
│
├── internal/
│   ├── actions/    # Workflow actions (transform, http_call)
│   ├── engine/     # Worker execution logic
│   ├── queue/      # Redis queue
│   ├── repository/ # Database access
│   ├── models/     # DB models
│   ├── config/     # Environment config
│   └── db/         # DB connection + migrations
│
├── migrations/     # SQL schema
├── docker-compose.yml
├── go.mod
└── README.md

⚙️ How to Run (Local)
1️⃣ Prerequisites

Go 1.21+

Docker + Docker Compose

Docker Desktop running (Windows / Mac)

2️⃣ Start Infrastructure
docker compose up -d


This starts:

PostgreSQL on 5432

Redis on 6379

3️⃣ Create .env file
DATABASE_URL=postgres://postgres:password@localhost:5432/workflow?sslmode=disable
REDIS_URL=redis://localhost:6379
PORT=8080

4️⃣ Run API Server (runs migrations)
go run cmd/api/main.go


Verify:

curl http://localhost:8080/health

5️⃣ Run Worker
go run cmd/worker/main.go


Worker runs continuously and waits for jobs.

▶️ Running a Sample Workflow
Insert Test Data (PostgreSQL)
INSERT INTO workflows (user_id, name, trigger_type)
VALUES (1, 'test-workflow', 'manual');

INSERT INTO workflow_steps (workflow_id, step_number, action_type, action_config)
VALUES (1, 1, 'transform', '{"select":["email"]}');

INSERT INTO trigger_events (workflow_id, payload)
VALUES (1, '{"email":"user@example.com","age":20}');

Push Job to Queue
go run push_job.go

Verify Execution
SELECT * FROM workflow_runs;
SELECT * FROM step_runs;


You should see:

workflow_run marked success

step_run with input/output logged

🔁 Supported Actions
transform

Selects specific fields from input JSON.

{ "select": ["email"] }

http_call

Makes external HTTP requests with template substitution.

{
  "url": "https://api.example.com",
  "method": "POST",
  "body_template": {
    "email": "{{email}}"
  }
}

🔐 Reliability Features

Per-step execution tracking

Retry logic (step-level)

Crash-safe execution

Explicit success/failure states

Timeouts for external calls

🎯 Why This Project Matters

This project demonstrates:

Real distributed systems thinking

Background processing

Durable state machines

Queue-based execution

Scalable backend architecture

This is the kind of system used inside workflow engines, schedulers, and automation platforms.

🛣️ Future Improvements

Webhook triggers

Cron scheduling

Idempotency keys

Execution dashboard APIs

Deployment to Railway / AWS

Horizontal worker scaling

👤 Author

Hemanth
Backend Engineer (Go, Distributed Systems)
