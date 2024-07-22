-- Add up migration script here
CREATE TABLE IF NOT EXISTS users (
    id TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    role TEXT NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS doctors (
    id TEXT UNIQUE NOT NULL,
    user_id TEXT NOT NULL,
    qualification TEXT,
    hospital TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS patients (
    id TEXT UNIQUE NOT NULL,
    user_id TEXT NOT NULL,
    doc_id TEXT NOT NULL,
    medical_history TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
