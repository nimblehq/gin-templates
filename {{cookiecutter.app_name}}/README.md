[![Build Status](CI_BADGE_URL goes here)](REPO_URL goes here)

## Introduction

> *App introduction goes here ...*

## Project Setup

### Prerequisites

- [Go - 1.16](https://golang.org/doc/go1.16) or newer

{% if cookiecutter.web_variant == "y" %}- [Node - 14](https://nodejs.org/en/){% endif %}

### Development

#### Create an ENV file

To start the development server, `.env` file must be created.

- Copy `.env.example` file and rename to `.env`

#### Build dependencies

- [`air`](https://github.com/cosmtrek/air) is used for live reloading

- [`goose`](https://github.com/pressly/goose) is used for database migration.

{% if cookiecutter.web_variant == "y" %}- [`forego`](https://github.com/ddollar/forego) manages Procfile-based applications.{% endif %}

They need to be built as a binary file in `$GOPATH`.


```make
make install-dependencies
```

#### Start development server

```make
make dev
```

The application runs locally at http://localhost:8080

### Test

Execute all unit tests:

```make
make test
```

### Migration

#### Create migration

```make
make migration/create MIGRATION_NAME={migration name}
```

#### List the migration status

```make
make migration/status
```

#### Migrate the database

```make
make db/migrate
```

#### Rollback the migration

```make
make db/rollback
```

{%- if cookiecutter.use_heroku == "yes" %}
### Deploy to Heroku with Terraform

#### Prerequisites

- [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli) latest version
- [Terraform](https://www.terraform.io/downloads.html)

To deploy the application to Heroku with Terraform, we need to create the Heroku API Key first:

```bash
$ heroku login
$ heroku authorizations:create --description <api key description>
```

And then, move to the `deploy/heroku` folder and run the following steps:

_Step 1:_ Copy the variable file and update the variables

```sh
$ cp terraform.tfvars.sample terraform.tfvars
```

*You can get the `tfvars` files from 1Password*

_Step 2:_ Initialize Terraform

```sh
$ terraform init
```

_Step 3:_ Generate an execution plan

```sh
$ terraform plan -var-file="terraform.tfvars"
```

_Step 5:_ Execute the generated plan

```sh
$ terraform apply -var-file="terraform.tfvars"
```

_Step 6:_ Build the application and push to heroku

You can check `.github/workflows/deploy.yml` workflow for more details

_Make sure you set the following Github secrets before deploying the application:_

```
HEROKU_ACCOUNT_EMAIL       # Heroku email
HEROKU_API_KEY             # Heroku OAuth token
HEROKU_APP_NAME_PRODUCTION # Heroku app name for production
HEROKU_APP_NAME_STAGING    # Heroku app name for staging
```
{%- endif %}
