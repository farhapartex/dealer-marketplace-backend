# Dealer Marketplace Backend

Go-based backend API using PostgreSQL and Redis.

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

Run tests
```bash
go test ./...
```
