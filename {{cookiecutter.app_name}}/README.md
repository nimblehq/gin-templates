[![Build Status](CI_BADGE_URL goes here)](REPO_URL goes here)

## Introduction

> *App introduction goes here ...*

## Project Setup

### Prerequisites

- [Go - 1.16](https://golang.org/doc/go1.16) or newer

### Development

#### Create an ENV file

To start development server, `.env` must be created.

- Copy `.env.example` and rename to `.env`

#### Build dependencies

- [`air`](https://github.com/cosmtrek/air) is used for live reloading

- [`goose`](https://github.com/pressly/goose) is used for database migration. 

They need to be built as a binary file in `$GOPATH`.


```sh
make install-dependencies
```

#### Start development server

```sh
make dev
```

The application runs locally at http://localhost:8080

### Test

Execute all unit tests:

```sh
make test
```
