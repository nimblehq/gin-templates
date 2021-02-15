.PHONY: env-setup env-teardown dev install-dependencies test

env-setup:
	docker-compose -f docker-compose.dev.yml up -d

env-teardown:
	docker-compose -f docker-compose.dev.yml down

dev: env-setup
	air -c cmd/api/.air.toml

install-dependencies:
	go get github.com/cosmtrek/air@v1.15.1
	go mod tidy

test:
	docker-compose -f docker-compose.test.yml up -d
	go test -v -p 1 -count=1 ./...
	docker-compose -f docker-compose.test.yml down
