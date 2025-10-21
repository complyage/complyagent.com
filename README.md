# ComplyAgents

ComplyAgents is a collection of **self-contained AI/ML micro-agents** designed for compliance and verification tasks.  
Each agent runs independently in its own Docker container and communicates exclusively through **RabbitMQ queues**.  
The system integrates with the [ComplyAge Go framework](https://github.com/ralphferrara/aria) for distributed task orchestration.

---

## ✨ Goals
- **Encapsulated services** — each agent includes its own Dockerfile, dependencies, and consumer logic.  
- **Queue-based communication** — all coordination occurs via RabbitMQ exchanges and queues.  
- **Lightweight** — no HTTP or gRPC interfaces; only internal queue-driven message flow.  
- **Docker-compose orchestration** — bring up all agents, RabbitMQ, and supporting services with one command.  
- **Extendable** — easily add new agents (OCR, DOB extraction, NSFW detection, Face compare, etc.).

---

## 📂 Repo Structure
```
complyagents/
├── maestro/        # RabbitMQ orchestrator / dispatcher
├── agents/         # Individual agent implementations
│   ├── ocr/
│   ├── dob/
│   ├── nsfw/
│   └── face/
├── docker-compose.yml
├── init.sh         # Linux init script
├── init.ps1        # PowerShell init script
└── README.md
```

---

## 🚀 Getting Started

### Prerequisites
- [Go 1.24+](https://go.dev/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop) with Compose v2
- (Optional) RabbitMQ locally or via Docker

### Cloning
```bash
git clone https://github.com/complyage/complyagent.com.git
cd complyagent.com
```

---

## 🔧 Usage

### Run an agent directly
Each agent can be started manually:
```bash
go run ./agents/ocr
```

### Run with docker-compose
```bash
docker compose up --build
```

This starts all defined agents, RabbitMQ, and the `maestro` orchestrator.

---

## 🧩 Planned Agents
- **OCR Agent** — text extraction from images  
- **DOB Agent** — date of birth recognition/validation  
- **NSFW Agent** — detect explicit/unsafe content  
- **Face Agent** — face comparison and selfie liveness  
- *(more to come…)*

---

## 🧠 Communication Model
- **RabbitMQ** handles all task distribution and message passing.  
- Each agent listens on its dedicated queue (e.g., `ocr.jobs`, `dob.jobs`, etc.).  
- Results are published back to response queues (e.g., `ocr.results`).  
- Messages are JSON objects defining task metadata, input paths, and output payloads.  

Example task flow:
1. The `maestro` service publishes a message to `dob.jobs`.
2. The DOB agent processes the job and publishes the result to `dob.results`.
3. Other services consume the result for downstream validation or reporting.

---

## 📜 License
MIT License © 2025
