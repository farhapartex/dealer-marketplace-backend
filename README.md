# Dealer Marketplace Backend

Go-based backend API using PostgreSQL and Redis.

![CI/CD Pipeline](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/ci.yml/badge.svg)

## Prerequisites

- Docker
- Docker Compose

## Installation

Clone the repository
```bash
cd dealer-market-be
```

Copy environment file
```bash
cp .env.docker .env
```

Update environment variables in .env file
```bash
nano .env
```

Build and start services
```bash
docker-compose up -d --build
```

View logs
```bash
docker-compose logs -f backend
```

Stop services
```bash
docker-compose down
```

Stop services and remove volumes
```bash
docker-compose down -v
```

## Get Database IP from Docker

Inspect database container
```bash
docker inspect dealer_postgres
```

Get database IP address
```bash
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' dealer_postgres
```

Get full network details
```bash
docker inspect dealer_postgres | grep IPAddress
```

List all containers with IPs
```bash
docker ps -q | xargs docker inspect -f '{{.Name}} - {{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'
```

## Run Database Migrations

Run a specific migration file
```bash
docker exec -i dealer_postgres psql -U devadmin -d dealer_db < migrations/001_create_users_table.sql
```

Run migration with connection from docker-compose
```bash
cat migrations/001_create_users_table.sql | docker exec -i dealer_postgres psql -U devadmin -d dealer_db
```

Connect to database shell
```bash
docker exec -it dealer_postgres psql -U devadmin -d dealer_db
```

Run all migrations in order
```bash
for file in migrations/*.sql; do docker exec -i dealer_postgres psql -U devadmin -d dealer_db < "$file"; done
```

## Endpoints

- API: http://localhost:8081
- PostgreSQL: localhost:5432
- Redis: localhost:6379

## Development

Build binary locally
```bash
go build -o bin/api
```

Run locally
```bash
./bin/api
```

## Testing

Test files are organized in the `tests/` directory by application module.

Run all tests
```bash
go test ./...
```

Run tests with verbose output
```bash
go test ./... -v
```

Run tests with coverage
```bash
go test ./... -cover
```

Run specific package tests
```bash
go test ./tests/user/services/... -v
```

Generate coverage report
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

Test directory structure
```
tests/
  user/
    services/
      auth_service_test.go
```

## CI/CD Pipeline

The project uses GitHub Actions for continuous integration and deployment.

Pipeline triggers
- Push to main branch
- Pull requests to main branch

Pipeline stages

**Build Stage:**
- Checkout code
- Setup Go environment
- Cache dependencies
- Install and verify dependencies
- Run go vet for code analysis
- Build application binary
- Build Docker image

**Test Stage:**
- Run all unit tests with race detection
- Generate coverage report
- Upload coverage artifacts

Workflow file location
```
.github/workflows/ci.yml
```

View pipeline status
- Check the Actions tab in your GitHub repository
- Badge status shown at the top of this README
