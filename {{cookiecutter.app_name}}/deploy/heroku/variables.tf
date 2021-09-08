variable "heroku_email" {
  description = "The email of Heroku account"
}

variable "heroku_api_key" {
  description = "The API key of Heroku account"
}

variable "heroku_organization" {
  description = "The team name of Heroku account"
}

variable "app_name" {
  description = "Heroku application name"
  default     = "{{cookiecutter.app_name}}"
}

variable "app_region" {
  description = "Heroku region"
  default     = "us"
}

variable "app_stack" {
  description = "Heroku stack"
  default     = "container"
}

variable "postgresql_plan" {
  description = "Heroku PostgreSQL addon plan"
  default     = "hobby-dev"
}

variable "dyno_size" {
  description = "Heroku Formation dyno size"
  default     = "Free"
}

variable "dyno_quantity" {
  description = "Heroku Formation dyno quantity"
  default     = 1
}
