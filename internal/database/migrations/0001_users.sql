CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY,
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now'))
) STRICT;

CREATE UNIQUE INDEX idx_users_email
ON users(email);
