CREATE TABLE users (
    id SERIAL PRIMARY KEY,

    library_external_id TEXT NOT NULL,
    name TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
)