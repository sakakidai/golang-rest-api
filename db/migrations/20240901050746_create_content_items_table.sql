-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS content_items(
   id bigserial PRIMARY KEY,
   name VARCHAR NOT NULL,
   description TEXT,
   created_at TIMESTAMP,
   updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS content_items;
-- +goose StatementEnd
