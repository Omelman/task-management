BIN_NAME := $(or $(PROJECT_NAME))
MIGRATE=migrate -path internal/sql/migrations -database postgres://postgres:12345@localhost:5432/manager?sslmode=disable

.PHONY: migrate-create migrate-up migrade-down

dep: ## Download required dependencies
	GO111MODULE=on go mod vendor

build: dep ## Build the binary file
	GO111MODULE=on go build -o ../bin/${BIN_NAME} -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/task-management

migrate-create: ## Create migration file with name
	migrate create -ext sql -dir internal/sql/migrations migration

db-up: ## Run db
	docker-compose up management_db

migrate-up: ## Run migrations
	$(MIGRATE) up

migrate-down: ## Rollback migrations
	$(MIGRATE) down

db-shell: ## Enter in db
	docker-compose -f docker-compose.yml exec management_db psql manager postgres