CREATE TABLE media_requests (
    id SERIAL PRIMARY KEY,

    requester_external_id TEXT NOT NULL,
    requested_by_user_id INTEGER NOT NULL REFERENCES users(id),
    available_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE INDEX media_requests_requested_by_user_id
ON media_requests (requested_by_user_id);