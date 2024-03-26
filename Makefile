include .env

migrate_up:
	migrate -path database/migrations -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" -verbose up

migrate_down:
	migrate -path database/migrations -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" -verbose down

migration:
	migrate create -ext sql -dir database/migrations -seq $(schema)

generate_docs:
	swag init --parseDependency --parseInternal

run_test:
	go test -v ./test

run:
	air

