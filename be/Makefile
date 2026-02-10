include .env
export

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

.PHONY: help run build db-up db-down migrate-up migrate-down migrate-create migrate-force

help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## Run the application
	go run ./cmd/api/

build: ## Build the application
	go build -o bin/api ./cmd/api/

db-up: ## Start PostgreSQL container
	docker compose up -d

db-down: ## Stop PostgreSQL container
	docker compose down

db-reset: ## Reset database (stop, remove volume, start)
	docker compose down -v
	docker compose up -d

migrate-up: ## Run all migrations
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migrate-down: ## Rollback all migrations
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

migrate-down-1: ## Rollback last migration
	migrate -path db/migrations -database "$(DB_URL)" -verbose down 1

migrate-create: ## Create a new migration (usage: make migrate-create name=migration_name)
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-force: ## Force migration version (usage: make migrate-force version=1)
	migrate -path db/migrations -database "$(DB_URL)" force $(version)
