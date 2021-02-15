.PHONY: dev install-dependencies

dev:
	air -c cmd/api/.air.toml

install-dependencies:
	go get github.com/cosmtrek/air@v1.15.1
	go mod tidy
