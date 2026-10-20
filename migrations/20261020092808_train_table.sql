-- +goose Up
CREATE TABLE IF NOT EXISTS trains(
    id TEXT PRIMARY KEY,
    name TEXT,
    origin TEXT,
    destination TEXT,
    origin_code TEXT,
    destination_code TEXT
);

-- +goose Down
DROP TABLE IF EXISTS trains;
