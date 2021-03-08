FROM golang:1.16-buster

ARG DATABASE_URL

ENV GIN_MODE=release

WORKDIR /app

# Install Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Install goose for migration
RUN go get github.com/pressly/goose/cmd/goose

# Copy codebase
COPY . .

# Run the migration
RUN goose -dir database/migrations -table "migration_versions" postgres "$DATABASE_URL" up

# Build go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/api

EXPOSE 8080

CMD ["./main"]
