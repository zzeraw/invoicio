COMPOSE = docker compose -f infra/docker-compose.yml --env-file infra/.env

.PHONY: up down lint migrate-up test fmt

up:
	$(COMPOSE) up -d --build

down:
	$(COMPOSE) down

lint:
	cd backend && golangci-lint run
	cd frontend && npm run lint

migrate-up:
	$(COMPOSE) run --rm migrate -path /migrations -database "postgres://$${POSTGRES_USER}:$${POSTGRES_PASSWORD}@db:5432/$${POSTGRES_DB}?sslmode=disable" up

test:
	cd backend && go test ./...
	cd frontend && npm run test

fmt:
	cd backend && gofmt -w ./cmd
	cd backend && goimports -w ./cmd
	cd frontend && npm run fmt
