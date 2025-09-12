# ComplyAgents

ComplyAgents is a collection of **self-contained AI/ML micro-agents** designed for compliance and verification tasks.  
Each agent exposes a **gRPC API** and runs in its own Docker container. The system is orchestrated with `docker-compose`  
and integrates with the [ComplyAge Go framework](https://github.com/ralphferrara/aria) for RabbitMQ task distribution.

---

## ✨ Goals
- **Encapsulated services** — each agent has its own Dockerfile, dependencies, and gRPC server.
- **Clear contracts** — all inter-service communication is defined in `.proto` files.
- **gRPC first** — no HTTP endpoints, only strongly-typed gRPC calls.
- **Docker-compose orchestration** — bring up all agents, RabbitMQ, and supporting services with a single command.
- **Extendable** — easily add new agents (OCR, DOB extraction, NSFW detection, Face compare, etc.).

---

## 📂 Repo Structure
```
complyagents/
├── proto/          # .proto service definitions
├── gen/            # generated Go code from protobufs
├── maestro/        # Consumer for RabbitMQ
├── agents/         # individual agent implementations
│   ├── ocr/
│   ├── dob/
│   ├── nsfw/
│   └── face/
├── docker/         # Dockerfiles and compose overrides
├── docker-compose.yml
├── init.sh         # Linux based init script
├── init.ps1        # Powershell init script
└── README.md
```

---

## 🚀 Getting Started

### Prerequisites
- [Go 1.24+](https://go.dev/)
- [Protocol Buffers compiler (`protoc`)](https://grpc.io/docs/protoc-installation/)
- Plugins for Go codegen:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```
- [Docker Desktop](https://www.docker.com/products/docker-desktop) with Compose v2

### Cloning
```bash
git clone https://github.com/<your-username>/complyagents.git
cd complyagents
```

---

## 🔧 Usage

### Generate gRPC stubs
```bash
protoc -I proto --go_out=gen --go-grpc_out=gen proto/*.proto
```

### Run locally
Each agent can be run directly:
```bash
go run ./agents/ocr
```

### Run with docker-compose
```bash
docker compose up --build
```

---

## 🧩 Planned Agents
- **OCR Agent** — text extraction from images
- **DOB Agent** — date of birth recognition/validation
- **NSFW Agent** — detect explicit/unsafe content
- **Face Agent** — face compare + selfie liveness
- (more to come…)

---

## 📜 License
MIT License © 2025
