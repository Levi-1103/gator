-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    CONSTRAINT user_id
    FOREIGN KEY(id)
    REFERENCES users(id)
);

-- +goose Down
DROP TABLE feeds;