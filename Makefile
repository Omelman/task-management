MIGRATE=migrate -path internal/sql/migrations -database postgres://postgres:12345@localhost:5432/manager?sslmode=disable

dep: ## Download required dependencies
	GO111MODULE=on go mod vendor

migrate-create: ## Create migration file with name
	migrate create -ext sql -dir internal/sql/migrations migration

migrate-up: ## Run migrations
	$(MIGRATE) up

migrate-down: ## Rollback migrations
	$(MIGRATE) down