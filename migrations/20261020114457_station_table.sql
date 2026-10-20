-- +goose Up
CREATE TABLE IF NOT EXISTS stations{
    id TEXT PRIMARY KEY AUTOINCREMENT,
    train_id TEXT
    station_code TEXT
    station_name TEXT
    time TEXT
}

-- +goose Down
DROP TABLE IF EXISTS stations
