-- Add up migration script here
CREATE TABLE sessions (
    user_id TEXT NOT NULL,
    session_id TEXT NOT NULL,
    expires_at TEXT NOT NULL
)
