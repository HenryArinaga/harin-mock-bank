CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email  VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL, 
    user_role VARCHAR(30) NOT NULL DEFAULT 'customer',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT users_role_check
        CHECK (user_role IN ('customer','admin','support'))
);

CREATE TABLE customers (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE REFERENCES users(id),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL, 
    phone VARCHAR(30),
    date_of_birth DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE accounts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES customers(id),
    account_status VARCHAR(30) NOT NULL DEFAULT 'pending',
    currency VARCHAR(3) NOT NULL,
    account_type VARCHAR(30) NOT NULL,
    account_number VARCHAR(32) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT accounts_status_check
        CHECK (account_status IN ('active', 'frozen', 'suspended', 'closed', 'pending')),

    CONSTRAINT accounts_type_check
        CHECK (account_type IN ('checking', 'savings', 'credit')),

    CONSTRAINT accounts_currency_check
        CHECK (LENGTH(currency) = 3 AND currency = UPPER(currency))

    
);

CREATE TABLE transactions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    to_account_id BIGINT REFERENCES accounts(id),
    from_account_id BIGINT REFERENCES accounts(id),
    transaction_type VARCHAR(30) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    transaction_description TEXT,
    transaction_status VARCHAR(30) NOT NULL DEFAULT 'pending',
    currency VARCHAR(3) NOT NULL,
    amount DECIMAL(19,4) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT transactions_type_check
        CHECK (transaction_type IN ('withdrawal', 'deposit', 'transfer')),

    CONSTRAINT transactions_status_check
        CHECK (transaction_status IN ('completed', 'pending', 'failed', 'cancelled', 'reversed')),

    CONSTRAINT transactions_currency_check
        CHECK (LENGTH(currency) = 3 AND currency = UPPER(currency)),

    CONSTRAINT transactions_amount_check
        CHECK (amount > 0),
    
    CONSTRAINT transactions_deposit_shape_check
        CHECK (transaction_type != 'deposit' OR (from_account_id IS NULL AND to_account_id IS NOT NULL)),

    CONSTRAINT  transactions_withdrawal_shape_check
        CHECK (transaction_type != 'withdrawal' OR (from_account_id IS NOT NULL AND to_account_id IS NULL)),

    CONSTRAINT transactions_transfer_shape_check
        CHECK (transaction_type != 'transfer' OR (from_account_id IS NOT NULL AND to_account_id IS NOT NULL)),

    CONSTRAINT transactions_transfer_distinct_accounts_check
        CHECK (transaction_type != 'transfer' OR from_account_id != to_account_id)

);

CREATE TABLE ledger_entries (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    account_id BIGINT NOT NULL REFERENCES accounts(id),
    transaction_id BIGINT NOT NULL REFERENCES transactions(id),
    direction VARCHAR(10) NOT NULL,
    amount DECIMAL(19,4) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT ledger_entries_currency_check
        CHECK (LENGTH(currency) = 3 AND currency = UPPER(currency)),

    CONSTRAINT ledger_entries_amount_check
        CHECK (amount > 0),
    
    CONSTRAINT ledger_entries_direction_check
        CHECK (direction IN ('debit', 'credit'))
    

);