CREATE TABLE IF NOT EXISTS products (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    stock int DEFAULT 0,
    price NUMERIC(10, 2) DEFAULT 0.00
);