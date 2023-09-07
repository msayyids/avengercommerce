CREATE TABLE IF NOT EXISTS transactions (
    id serial PRIMARY KEY,
    user_id int,
    product_id int,
    quantity int NOT NULL,
    total_amount int NOT null,
    store_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (product_id) REFERENCES products (id),
    FOREIGN KEY (store_id) REFERENCES store(id),
);