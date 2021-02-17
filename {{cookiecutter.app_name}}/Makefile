.PHONY: dev install-dependencies test

dev:
	air -c cmd/api/.air.toml

install-dependencies:
	go get github.com/cosmtrek/air@v1.15.1
	go mod tidy

test:
	go test -v -p 1 -count=1 ./...
