# Workflow Automation Engine (Zapier-like Backend)

A production-style **workflow automation engine** built in Go that executes multi-step workflows asynchronously using Redis queues and PostgreSQL for durable state tracking.

This project demonstrates how systems like **Zapier, n8n, Temporal, GitHub Actions** work internally.

---

## Features

- Asynchronous workflow execution using background workers  
- Multi-step workflows with input → output chaining  
- Durable execution state with workflow runs and step runs  
- Redis-backed job queue for scalability and concurrency  
- PostgreSQL for reliability, retries, and observability  
- Clean separation between API server and worker engine  

This is **not a CRUD application** — it is a backend orchestration system.

---

## Architecture Overview

API Server (Gin)
├─ stores workflows and steps
├─ receives trigger events
└─ pushes jobs to Redis

Redis Queue
└─ holds workflow execution jobs

Worker (Go)
├─ runs continuously
├─ pulls jobs from Redis
├─ executes workflow steps sequentially
└─ persists execution state to PostgreSQL

yaml
Copy code

---

## Tech Stack

- **Language**: Go  
- **HTTP Framework**: Gin  
- **Database**: PostgreSQL (sqlx)  
- **Queue**: Redis  
- **Infrastructure**: Docker, Docker Compose  

---

## Project Structure

workflow-engine/
├── cmd/
│ ├── api/ # API server
│ └── worker/ # Background worker
│
├── internal/
│ ├── actions/ # Workflow actions (transform, http_call)
│ ├── engine/ # Worker execution logic
│ ├── queue/ # Redis queue
│ ├── repository/ # Database access layer
│ ├── models/ # Data models
│ ├── config/ # Environment configuration
│ └── db/ # DB connection and migrations
│
├── migrations/ # SQL schema
├── docker-compose.yml
├── go.mod
└── README.md

yaml
Copy code

---

## Getting Started

### Prerequisites

- Go 1.21 or newer  
- Docker & Docker Compose  
- Docker Desktop running (Windows / macOS)

---

### Start Infrastructure

```bash
docker compose up -d
This starts:

PostgreSQL on port 5432

Redis on port 6379

Create .env File
Create a .env file in the project root:

env
Copy code
DATABASE_URL=postgres://postgres:password@localhost:5432/workflow?sslmode=disable
REDIS_URL=redis://localhost:6379
PORT=8080
Run API Server
bash
Copy code
go run cmd/api/main.go
This:

connects to PostgreSQL

runs migrations

starts a minimal HTTP server

Verify:

bash
Copy code
curl http://localhost:8080/health
Run Worker
bash
Copy code
go run cmd/worker/main.go
The worker runs continuously and waits for jobs from Redis.

Running a Sample Workflow
Insert Test Data (PostgreSQL)
sql
Copy code
INSERT INTO workflows (user_id, name, trigger_type)
VALUES (1, 'test-workflow', 'manual');

INSERT INTO workflow_steps (workflow_id, step_number, action_type, action_config)
VALUES (1, 1, 'transform', '{"select":["email"]}');

INSERT INTO trigger_events (workflow_id, payload)
VALUES (1, '{"email":"user@example.com","age":20}');
Push Job to Redis
bash
Copy code
go run push_job.go
Verify Execution
sql
Copy code
SELECT id, status FROM workflow_runs;
SELECT input_data, output_data FROM step_runs;
You should see:

workflow run marked as success

step output containing only the selected fields

Supported Actions
Transform
Selects specific fields from input JSON.

json
Copy code
{ "select": ["email"] }
HTTP Call
Makes an external HTTP request with template substitution.

json
Copy code
{
  "url": "https://api.example.com",
  "method": "POST",
  "body_template": {
    "email": "{{email}}"
  }
}
Reliability Guarantees
Per-step execution tracking

Step-level retry logic

Crash-safe execution

Explicit workflow success/failure states

Timeouts on external HTTP calls

Why This Project
This project demonstrates:

Background job processing

Queue-based architectures

Durable state machines

Workflow orchestration design

Scalable backend engineering

It reflects real systems engineering, not tutorial-level CRUD work.

Future Improvements
Webhook triggers

Cron scheduling

Idempotency keys

Execution history APIs

Deployment to Railway or AWS

Horizontal worker scaling

Author
Hemanth
Backend Engineer (Go, Distributed Systems)
