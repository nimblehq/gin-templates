-- +goose Up
-- +goose StatementBegin
CREATE TABLE healths (
  id SERIAL PRIMARY KEY,
  status TEXT UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE healths;
-- +goose StatementEnd
