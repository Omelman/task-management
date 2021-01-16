BIN_NAME := $(or $(PROJECT_NAME),'api')

MIGRATE=migrate -path api/sql/migrations -database postgres://postgres:12345@localhost:5432/manager?sslmode=disable

dep: ## Download required dependencies
	GO111MODULE=on go mod vendor

build: dep ## Build the binary file
	GO111MODULE=on go build -o ../bin/${BIN_NAME} -a -tags netgo -ldflags '-w -extldflags "-static"' .

migrate-create: ## Create migration file with name
	migrate create -ext sql -dir api/sql/migrations migration

migrate-up: ## Run migrations
	$(MIGRATE) up

migrate-down: ## Rollback migrations
	$(MIGRATE) down