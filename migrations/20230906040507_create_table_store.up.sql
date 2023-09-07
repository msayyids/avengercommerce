CREATE TABLE IF NOT EXISTS stores (
    id serial PRIMARY KEY,
    store_name VARCHAR(255) NOT NULL,
    rating INT CHECK (rating >= 0 AND rating <= 10) DEFAULT 0,
    longitude FLOAT NOT NULL,
    latitude FLOAT NOT NULL,
    address VARCHAR(255) NOT NULL
);
