CREATE TABLE inbounds (
    id SERIAL PRIMARY KEY,
    created_by INT,
    note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE inbound_items (
    id SERIAL PRIMARY KEY,
    inbound_id INT REFERENCES inbounds(id) ON DELETE CASCADE,
    item_id INT REFERENCES items(id),
    qty INT NOT NULL,
    price_buy NUMERIC(12,2)
);