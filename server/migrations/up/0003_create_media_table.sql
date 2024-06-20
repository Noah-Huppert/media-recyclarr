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
    type MEDIATYPE NOT NULL,
    node_path LTREE NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE INDEX media_node_path ON media USING GIST (node_path);