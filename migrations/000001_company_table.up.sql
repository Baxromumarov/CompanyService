CREATE TABLE companies (
    id UUID PRIMARY KEY,
    name VARCHAR(15) UNIQUE NOT NULL,
    description TEXT,
    amount_of_employees INT NOT NULL,
    registered BOOLEAN NOT NULL,
    type VARCHAR(50) CHECK (
        type IN (
            'Corporations',
            'NonProfit',
            'Cooperative',
            'Sole Proprietorship'
        )
    ) NOT NULL
);