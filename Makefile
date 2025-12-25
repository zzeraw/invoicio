COMPOSE = docker compose -f infra/docker-compose.yml --env-file infra/.env

.PHONY: up down lint test fmt

up:
	$(COMPOSE) up -d --build

down:
	$(COMPOSE) down

lint:
	cd backend && golangci-lint run
	cd frontend && npm run lint

test:
	cd backend && go test ./...
	cd frontend && npm run test

fmt:
	cd backend && gofmt -w ./cmd
	cd backend && goimports -w ./cmd
	cd frontend && npm run fmt
