CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS Wallets(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    balance FLOAT CHECK (balance > 0 )
);

CREATE TABLE transactions (
    id UUID DEFAULT uuid_generate_v4(),
    time TIMESTAMP,
    from_wallet UUID,
    to_wallet UUID,
    amount FLOAT,
    CONSTRAINT fk_wallets_to FOREIGN KEY(to_wallet) REFERENCES Wallets(id),
    CONSTRAINT fk_wallets_from FOREIGN KEY(from_wallet) REFERENCES Wallets(id)
);

