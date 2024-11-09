-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pg_uuidv7;

CREATE TABLE team (
    id text PRIMARY KEY DEFAULT uuid_generate_v7 (),
    team_name text NOT NULL,
    secret text NOT NULL
);

CREATE TABLE flag (
    id text PRIMARY KEY DEFAULT uuid_generate_v7 (),
    question text NOT NULL,
    answer text NOT NULL,
    answer_option jsonb NOT NULL,
    score integer NOT NULL,
    seq_num integer NOT NULL,
    visited boolean,
    correct boolean,
    team_id text NOT NULL REFERENCES team,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE flag;
DROP TABLE team;
-- +goose StatementEnd
