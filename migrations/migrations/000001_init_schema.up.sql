CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    sku VARCHAR(100) UNIQUE,
    barcode VARCHAR(100),
    price_sell NUMERIC(12,2) NOT NULL,
    price_buy NUMERIC(12,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE stocks (
    id SERIAL PRIMARY KEY,
    item_id INT UNIQUE REFERENCES items(id) ON DELETE CASCADE,
    quantity INT DEFAULT 0,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE stock_movements (
    id SERIAL PRIMARY KEY,
    item_id INT REFERENCES items(id),
    type VARCHAR(20),
    qty INT,
    reference_type VARCHAR(50),
    reference_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);