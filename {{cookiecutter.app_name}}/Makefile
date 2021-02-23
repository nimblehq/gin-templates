.PHONY: env-setup env-teardown dev install-dependencies test

env-setup:
	docker-compose -f docker-compose.dev.yml up -d

env-teardown:
	docker-compose -f docker-compose.dev.yml down

db/migrate:
	goose -dir database/migrations -table "migration_versions" postgres "user=postgres dbname=xxxx_development sslmode=disable" up

db/rollback:
	goose -dir database/migrations -table "migration_versions" postgres "user=postgres dbname=xxxx_development sslmode=disable" down

migration/create:
ifndef MIGRATION_NAME
	$(error MIGRATION_NAME is required)
endif
	goose -dir database/migrations create ${MIGRATION_NAME} sql

migration/status:
	goose -dir database/migrations -table "migration_versions" postgres "user=postgres dbname=xxxx_development sslmode=disable" status

dev: env-setup
	air -c cmd/api/.air.toml

install-dependencies:
	go get github.com/cosmtrek/air@v1.15.1
	go get github.com/pressly/goose/cmd/goose
	go mod tidy

test:
	docker-compose -f docker-compose.test.yml up -d
	BRANCH=$(BRANCH) go test -v -p 1 -count=1 ./...
	docker-compose -f docker-compose.test.yml down
