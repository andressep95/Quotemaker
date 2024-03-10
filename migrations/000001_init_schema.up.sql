-- Table category
CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL
);

-- Table product
CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id INT,
    length FLOAT, 
    price FLOAT,
    weight FLOAT,
    code VARCHAR(255),
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
    delivered_at TIMESTAMP,  -- Fecha de entrega
    FOREIGN KEY (seller_id) REFERENCES seller(id),
    FOREIGN KEY (customer_id) REFERENCES customer(id)
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
