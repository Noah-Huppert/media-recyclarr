CREATE TABLE user_media_finishes (
    id SERIAL PRIMARY KEY,
    
    user_id INTEGER NOT NULL REFERENCES users(id),
    media_request_id INTEGER NOT NULL REFERENCES media_requests(id),
    
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE INDEX user_media_finishes_user_id
ON user_media_finishes (user_id);

CREATE INDEX user_media_finishes_media_request_id
ON user_media_finishes (media_request_id);