
CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY,
    owner_id UUID REFERENCES customers (id) ON DELETE RESTRICT,
    balance BIGINT CHECK(balance >= 0) DEFAULT 0,
    currency_code VARCHAR(255) NOT NULL DEFAULT 'UZS',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);


CREATE INDEX idx_owner_id ON accounts (owner_id);

