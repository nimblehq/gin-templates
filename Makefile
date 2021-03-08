.PHONY: test

test:
	BRANCH=$(BRANCH) go test -v -p 1 -count=1 ./...
