CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    deposit_amount NUMERIC(10, 2) DEFAULT 0.00
);
