CREATE TABLE sales (
    id SERIAL PRIMARY KEY,
    customer_name VARCHAR(150),
    total_amount NUMERIC(12,2),
    created_by INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sales_items (
    id SERIAL PRIMARY KEY,
    sales_id INT REFERENCES sales(id) ON DELETE CASCADE,
    item_id INT REFERENCES items(id),
    qty INT NOT NULL,
    price_sell NUMERIC(12,2)
);