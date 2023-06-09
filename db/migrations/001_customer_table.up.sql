CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY,
    phone_number VARCHAR(16) UNIQUE,
    name VARCHAR(64),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

