.PHONY: run
## Run server. Usage: 'make run'
run: ; $(info running server...) @
	@$(DEV_VARS) go run ./cmd/server/main.go

.PHONY: migrations
## Run migrations. Usage: 'make migrations'
migrations: ; $(info running migrations...) @
	@$(DEV_VARS) go run ./cmd/migrations/main.go