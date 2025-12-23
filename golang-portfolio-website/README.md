# Portfolio (Golang)

This is a small Golang portfolio web application (server-side rendered templates + JSON API) built with chi, html/template, pgx/Postgres, and simple service/repository layers.

## Prerequisites
- Go 1.20+ installed (go toolchain)
- (Optional for running server) PostgreSQL database
- (Optional) Docker to quickly run Postgres during development

## Quick start (dev)

1. Clone repository

    git clone <your-repo-url>
    cd golang-portfolio-website

2. (Optional) Run a PostgreSQL container for local development

    docker run --name pf-db -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=portfolio -p 5432:5432 -d postgres:15

3. Set environment variables

    export DATABASE_URL="postgres://postgres:pass@localhost:5432/portfolio"
    export SERVER_PORT=8080

4. Run the app

    go run ./cmd/server

The server will listen on `:8080` (or `SERVER_PORT` if set).

Notes:
- The server expects a database with the required tables; if you have SQL migrations or provisioning steps, run them before starting the server.

## Tests

Run unit tests:

    go test ./... -v

Run tests with coverage:

    go test ./... -coverprofile=coverage.out
    go tool cover -func=coverage.out

Coverage is enforced in CI (see `.github/workflows/ci.yml`).

## Project structure (high-level)

- `cmd/server` - app entrypoint
- `internal/handler/http` - HTTP handlers
- `internal/service` - business logic
- `internal/repository/postgres` - Postgres repository implementations
- `templates/` - HTML templates
- `internal/util` - helpers (logger, validator, response helpers)

## CI
The project includes a GitHub Actions workflow that runs tests and checks coverage on `push`/`pull_request`.

If you'd like, I can also add a `docker-compose.yml` and SQL migrations to simplify setup further.

---
Happy hacking! If you want, I can now add the CI workflow that fails when coverage < 50%.
# Portfolio Website - Golang

A professional portfolio website built with Golang following clean architecture principles.

## Tech Stack

- **Backend**: Go 1.21+
- **Web Framework**: chi router
- **Database**: PostgreSQL with pgx/v5 driver
- **Templates**: HTML templates
- **Styling**: Bootstrap 5
- **Logger**: Zap
- **Validation**: go-playground/validator

## Project Structure

```
project-app-portfolio-golang-rafli/
├── cmd/server/          # Application entry point
├── internal/            # Private application code
│   ├── handler/         # HTTP handlers
│   ├── service/         # Business logic
│   ├── repository/      # Data access layer
│   ├── model/           # Domain models
│   ├── dto/             # Data transfer objects
│   ├── util/            # Utilities
│   ├── config/          # Configuration
│   └── router/          # Route definitions
├── templates/           # HTML templates
├── static/              # Static assets
├── migrations/          # Database migrations
└── tests/              # Test files
```

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Setup PostgreSQL database:
```bash
createdb portfolio_db
```

3. Run migrations:
```bash
psql -d portfolio_db -f migrations/001_create_users.sql
psql -d portfolio_db -f migrations/002_create_projects.sql
psql -d portfolio_db -f migrations/003_create_contacts.sql
```

4. Configure environment variables:
```bash
cp .env.example .env
# Edit .env with your database credentials
```

5. Run the application:
```bash
go run cmd/server/main.go
```

6. Visit http://localhost:8080

## Features

- Clean architecture with handler → service → repository layers
- RESTful API endpoints
- Server-side rendered HTML templates
- PostgreSQL database with pgx/v5
- Input validation
- Structured logging with Zap
- Responsive design with Bootstrap 5
- Contact form with AJAX submission
- Project showcase
- Professional design

## API Endpoints

- `GET /` - Homepage
- `GET /projects` - Projects list
- `GET /contact` - Contact form
- `POST /api/contact` - Submit contact form
- `GET /api/projects` - Get all projects (JSON)
- `POST /api/projects` - Create project (JSON)
- `GET /api/projects/{id}` - Get project by ID (JSON)
- `GET /health` - Health check

## Testing

Run tests with:
```bash
go test ./tests/... -v
```

## License

MIT License
```

```makefile file="Makefile"
.PHONY: run build test clean migrate

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

test:
	go test -v ./tests/...

clean:
	rm -rf bin/

migrate:
	psql -d portfolio_db -f migrations/001_create_users.sql
	psql -d portfolio_db -f migrations/002_create_projects.sql
	psql -d portfolio_db -f migrations/003_create_contacts.sql

deps:
	go mod download
	go mod tidy
