-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS Clients
(
        id        BIGSERIAL PRIMARY KEY NOT NULL,
        client_name TEXT,
        version INT,
        image   TEXT,
        cpu     TEXT,
        memory  TEXT,
        priority FLOAT,
        need_restart BOOLEAN,
        spawned_at timestamptz DEFAULT NOW(),
        created_at timestamptz DEFAULT NOW(),
        updated_at timestamptz DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS AlgorithmStatus
(
    id        BIGSERIAL PRIMARY KEY NOT NULL,
    client_id BIGINT                NOT NULL references Clients(id),
    VWAP      BOOLEAN               NOT NULL DEFAULT FALSE,
    TWAP      BOOLEAN               NOT NULL DEFAULT FALSE,
    HFT       BOOLEAN               NOT NULL DEFAULT FALSE
);

CREATE UNIQUE INDEX algorithm_status_unique_index ON AlgorithmStatus(client_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS AlgorithmStatus;
DROP TABLE IF EXISTS Clients;
-- +goose StatementEnd
