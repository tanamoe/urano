-- +goose Up
CREATE TABLE registry (
    id uuid PRIMARY KEY DEFAULT uuidv7(),
    registration_id text UNIQUE NOT NULL,
    isbn text,
    title text NOT NULL,
    author text,
    translator text,
    print_amount integer,
    self_publish boolean,
    partner text
);

CREATE INDEX registry_registration_id_index ON registry (registration_id);

-- +goose Down
DROP TABLE registry;
