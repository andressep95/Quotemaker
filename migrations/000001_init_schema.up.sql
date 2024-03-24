-- Table category
CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL
);

-- Table product
CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    category_id INT,
    code VARCHAR(255),
    description VARCHAR(255) NOT NULL,
    price FLOAT,
    weight FLOAT,
    length FLOAT, 
    is_available BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (category_id) REFERENCES category(id)
);

-- Table quotation
CREATE TABLE IF NOT EXISTS quotation (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    total_price FLOAT,
    is_purchased BOOLEAN DEFAULT FALSE,
    purchased_at TIMESTAMP,  -- Fecha de compra
    is_delivered BOOLEAN DEFAULT FALSE,
    delivered_at TIMESTAMP  -- Fecha de entrega
);

-- Table quote_product
CREATE TABLE IF NOT EXISTS quote_product (
    id SERIAL PRIMARY KEY,
    quotation_id INT,
    product_id INT,
    quantity FLOAT,
    FOREIGN KEY (quotation_id) REFERENCES quotation(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
);


-- Indexes
CREATE INDEX idx_category_name ON category (category_name);

CREATE INDEX idx_product_name ON product (description);
CREATE INDEX idx_category_id ON product (category_id);
CREATE INDEX idx_product_code ON product (code);
CREATE INDEX idx_product_price ON product (price);

CREATE INDEX idx_quotation_created_at ON quotation (created_at);
CREATE INDEX idx_quotation_is_purchased ON quotation (is_purchased);
CREATE INDEX idx_quotation_is_delivered ON quotation (is_delivered);

CREATE INDEX idx_quote_product_quotation_id ON quote_product (quotation_id);
CREATE INDEX idx_quote_product_product_id ON quote_product (product_id);
