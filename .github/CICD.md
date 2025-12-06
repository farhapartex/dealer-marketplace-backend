# CI/CD Pipeline Documentation

## Overview

This project uses GitHub Actions for continuous integration and continuous deployment. The pipeline automatically runs on every push to the main branch and on pull requests targeting main.

## Workflow File

Location: `.github/workflows/ci.yml`

## Pipeline Stages

### Stage 1: Build and Docker Check

**Purpose:** Validate code compilation and Docker image build

**Steps:**
1. Checkout repository code
2. Setup Go 1.25.4 environment
3. Cache Go modules for faster builds
4. Download and verify Go dependencies
5. Run `go vet` for static analysis
6. Build application binary
7. Verify binary was created successfully
8. Setup Docker Buildx
9. Build Docker image (without pushing to registry)
10. Verify Docker build success

**Exit Conditions:**
- Fails if Go build errors occur
- Fails if binary is not created
- Fails if Docker build fails

### Stage 2: Run Unit Tests

**Purpose:** Execute all unit tests and generate coverage report

**Dependencies:** Requires Build stage to pass first

**Steps:**
1. Checkout repository code
2. Setup Go 1.25.4 environment
3. Cache Go modules for faster builds
4. Download Go dependencies
5. Run all tests with race detection enabled
6. Generate coverage report
7. Display total coverage percentage
8. Upload coverage report as artifact

**Exit Conditions:**
- Fails if any test fails
- Fails if race conditions detected

## Triggers

### Push to Main
```yaml
on:
  push:
    branches:
      - main
```

### Pull Request to Main
```yaml
on:
  pull_request:
    branches:
      - main
```

## Artifacts

**Coverage Report:**
- Name: coverage-report
- File: coverage.out
- Retention: 7 days
- Download from: Actions > Workflow Run > Artifacts section

## Local Testing

Test build stage locally
```bash
go mod download
go mod verify
go vet ./...
go build -v -o bin/api .
docker build -t dealer-market-be:local .
```

Test test stage locally
```bash
go test ./... -v -race -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## Performance Optimizations

**Caching:**
- Go modules cached using go.sum hash
- Docker layers cached using GitHub Actions cache
- Reduces build time significantly on subsequent runs

## Status Badge

Add to README.md
```markdown
![CI/CD Pipeline](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/ci.yml/badge.svg)
```

Replace YOUR_USERNAME and YOUR_REPO with actual values

## Troubleshooting

**Build fails with "go version mismatch"**
- Ensure go.mod specifies go 1.25.4
- Update workflow file if Go version changes

**Tests fail in CI but pass locally**
- Check for race conditions: `go test ./... -race`
- Verify dependencies: `go mod verify`
- Check environment-specific code

**Docker build fails**
- Verify Dockerfile syntax
- Test locally: `docker build -t test .`
- Check .dockerignore is not excluding required files

## Future Enhancements

Planned additions:
- Docker image push to registry (Docker Hub, ECR, etc.)
- Deployment to staging/production environments
- Code quality checks (golangci-lint)
- Security scanning (gosec)
- Integration tests
- Performance benchmarks
