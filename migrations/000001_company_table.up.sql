CREATE TYPE company_type AS ENUM (
    'Corporations',
    'NonProfit',
    'Cooperative',
    'Sole Proprietorship'
);

CREATE TABLE companies (
    id UUID PRIMARY KEY,
    name VARCHAR(15) UNIQUE NOT NULL,
    description TEXT,
    amount_of_employees INT NOT NULL,
    registered BOOLEAN NOT NULL,
    type company_type NOT NULL
);