include .env
ifdef ENV
include .env.$(ENV)
endif

.PHONY: env-setup env-teardown db/migrate db/rollback migration/create migration/status dev install-dependencies test wait-for-postgres

env-setup:
	docker compose -f docker-compose.dev.yml up -d

env-teardown:
	docker compose -f docker-compose.dev.yml down

db/migrate:
	make wait-for-postgres
	goose -dir database/migrations -table "migration_versions" postgres "$(DATABASE_URL)" up

db/rollback:
	make wait-for-postgres
	goose -dir database/migrations -table "migration_versions" postgres "$(DATABASE_URL)" down

migration/create:
ifndef MIGRATION_NAME
	$(error MIGRATION_NAME is required)
endif
	goose -dir database/migrations create $(MIGRATION_NAME) sql

migration/status:
	goose -dir database/migrations -table "migration_versions" postgres "$(DATABASE_URL)" status

{% if cookiecutter._api_variant == "yes" %}doc/generate:
	npm run build:docs
{%- endif %}

dev:
	make env-setup
	make db/migrate
	forego start -f Procfile.dev

install-dependencies:
	go install github.com/cosmtrek/air@v1.43.0
	go install github.com/pressly/goose/v3/cmd/goose@v3.9.0
	go install github.com/ddollar/forego@latest
	go mod tidy
	npm install
	{% if cookiecutter._api_variant == "yes" %}make doc/generate{%- endif %}

test:
	docker compose -f docker-compose.test.yml up -d
	make wait-for-postgres
	go test -v -p 1 -count=1 ./... -coverprofile=coverage/coverage.out -ginkgo.reportFile=coverage/test-report.xml
	docker compose -f docker-compose.test.yml down

coverage:
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	go tool cover -func=coverage/coverage.out

coverage-check: coverage
	./bin/check-coverage.sh

wait-for-postgres:
	$(shell DATABASE_URL=$(DATABASE_URL) ./bin/wait-for-postgres.sh)
