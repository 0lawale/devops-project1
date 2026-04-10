# DevOps Project 1 — Go Task Manager API with GitHub Actions

A production-grade REST API built in Go, fully containerised with Docker and deployed
automatically to AWS EC2 via a GitLab CI/CD pipeline.

---


## Project Structure

```bash
.
├── go.mod
├── go.sum
├── handlers.go
├── main.go
├── main_test.go
└── README.md

1 directory, 6 files
```
---

---

## Tech Stack

| Tool | Purpose |
|------|---------|
| Go 1.23 | Application language |
| Docker | Containerisation (multi-stage build, ~15MB image) |
| GitHub Actions | Automated pipeline |
| Docker Hub | Container image registry |
| AWS EC2 | Deployment target |
| golangci-lint | Static code analysis |

---

## CI/CD Pipeline Stages

| Stage | What it does |
|-------|-------------|
| **lint** | Runs `golangci-lint` to enforce code quality |
| **test** | Runs all unit tests with coverage report |
| **build** | Compiles a static Go binary |
| **dockerise** | Builds Docker image, tags with commit SHA, pushes to Docker Hub |
| **deploy** | SSHs into EC2, pulls latest image and runs container |

Every push to `main` triggers the full pipeline automatically.

---

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/tasks` | Get all tasks |
| GET | `/tasks/{id}` | Get task by ID |
| POST | `/tasks` | Create a new task |
| DELETE | `/tasks/{id}` | Delete a task |

---

## Running Locally

**Prerequisites:** Go 1.23+, Docker

```bash
# Clone the repo
git clone https://github.com/0lawale/devops-project1.git
cd devops-project1

# Run directly
go run .

# Or with Docker
docker build -t devops-project1 .
docker run -p 8080:8080 devops-project1
```

Test it:
```bash
# Health check
curl http://localhost:8080/health

# Create a task
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"My first task"}'

# Get all tasks
curl http://localhost:8080/tasks

# Delete a task
curl -X DELETE http://localhost:8080/tasks/1
```

---

## Project Highlights

- **Multi-stage Docker build** — final image is ~15MB by copying only the compiled binary into a minimal Alpine base
- **Non-root container** — app runs as an unprivileged user for security
- **GitHub CI/CD** — full pipeline from lint to live deployment on every push to main
- **Automated deployment** — zero manual steps to get code to production
- **Commit-tagged images** — every Docker image is tagged with the Git commit SHA for full traceability

---

## CI/CD Variables Required

| Variable | Description |
|----------|-------------|
| `EC2_HOST` | Public IP of the EC2 instance |
| `EC2_USER` | SSH user (`ubuntu`) |
| `EC2_SSH_KEY` | Base64 encoded private key |
| `DOCKER_HUB_USERNAME` | Docker Hub username |
| `DOCKER_HUB_TOKEN` | Docker Hub access token |
