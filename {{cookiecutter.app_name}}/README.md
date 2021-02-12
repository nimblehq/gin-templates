[![Build Status](CI_BADGE_URL goes here)](REPO_URL goes here)

## Introduction

> *App introduction goes here ...*

## Project Setup

### Prerequisites

- [Go - 1.15](https://golang.org/doc/go1.15) or newer

### Development

#### Build dependencies
[`air`](https://github.com/cosmtrek/air) is used for live reloading. It needs to be built as a binary file in `$GOPATH`.


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
