CREATE TYPE MEDIATYPE AS ENUM (
    'movie',
    'series',
    'season',
    'episode'
);

CREATE TABLE media (
    id SERIAL PRIMARY KEY,

    library_external_id TEXT NOT NULL,
    name TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);