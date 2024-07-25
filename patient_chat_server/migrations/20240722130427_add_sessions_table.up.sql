-- Add up migration script here
CREATE TABLE sessions (
    user_id TEXT NOT NULL,
    session_id TEXT NOT NULL PRIMARY KEY,
    expires_at TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
)
