
CREATE TABLE transaction_types (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

INSERT INTO transaction_types (id, name, created_at) 
VALUES (0, 'income', CURRENT_TIMESTAMP), 
        (1, 'expense', CURRENT_TIMESTAMP),
        (2, 'transfer', CURRENT_TIMESTAMP);

--('topup', 'withdraw', 'transfer', 'debt');



CREATE TABLE IF NOT EXISTS transactions(
    id UUID PRIMARY KEY,
    type_id INTEGER NOT NULL REFERENCES transaction_types(id),--topup, withdraw, transfer
    account_from_id UUID REFERENCES accounts(id) ON DELETE RESTRICT ,
    account_to_id UUID REFERENCES accounts(id) ON DELETE RESTRICT,
    amount BIGINT NOT NULL CHECK (amount > 0),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE transactions
    ADD CONSTRAINT account_to_id_check  
        CHECK (((type_id=1/*'expense'*/) AND (account_to_id IS NULL)) 
            OR ((type_id<>1/*'not expense'*/) AND (account_to_id IS NOT NULL))),
    ADD CONSTRAINT account_from_id_check 
      CHECK(((type_id=0/*'income'*/) AND (account_from_id IS NULL)) 
            OR ((type_id<>0/*'not income'*/) AND (account_from_id IS NOT NULL))) ;



CREATE INDEX idx_account_from_id ON transactions(account_from_id);
CREATE INDEX idx_account_to_id ON transactions(account_to_id);