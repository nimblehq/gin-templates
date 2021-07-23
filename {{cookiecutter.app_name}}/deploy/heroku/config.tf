resource "heroku_config" "common" {
  # Environment variables
  vars = {
    GIN_MODE = "release"
  }
}

resource "heroku_app_config_association" "default" {
  app_id = heroku_app.default.id

  vars = heroku_config.common.vars
}
