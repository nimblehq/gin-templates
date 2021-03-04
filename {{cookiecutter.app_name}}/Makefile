.PHONY: env-setup env-teardown db/migrate db/rollback migratoin/create migration/status dev install-dependencies test

env-setup:
	docker-compose -f docker-compose.dev.yml up -d

env-teardown:
	docker-compose -f docker-compose.dev.yml down

db/migrate:
	make check-env
	goose -dir database/migrations -table "migration_versions" postgres "postgres://postgres:postgres@0.0.0.0:${DB_PORT}/test-gin-template_${ENV}?sslmode=disable" up

db/rollback:
	goose -dir database/migrations -table "migration_versions" postgres "postgres://postgres:postgres@0.0.0.0:$(DB_PORT)/test-gin-template_${ENV}?sslmode=disable" down

migration/create:
ifndef MIGRATION_NAME
	$(error MIGRATION_NAME is required)
endif
	goose -dir database/migrations create ${MIGRATION_NAME} sql

migration/status:
	goose -dir database/migrations -table "migration_versions" postgres "user=postgres dbname=test-gin-template_${ENV} sslmode=disable" status

dev: 
	make env-setup
	ENV=development make db/migrate
	air -c cmd/api/.air.toml

install-dependencies:
	go get github.com/cosmtrek/air@v1.15.1
	go get github.com/pressly/goose/cmd/goose
	go mod tidy

test:
	docker-compose -f docker-compose.test.yml up -d
	ENV=test make db/migrate
	BRANCH=$(BRANCH) go test -v -p 1 -count=1 ./...
	docker-compose -f docker-compose.test.yml down
