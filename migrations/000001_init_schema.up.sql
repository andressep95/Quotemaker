-- Table seller
CREATE TABLE IF NOT EXISTS seller (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Table customer
CREATE TABLE IF NOT EXISTS customer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    rut VARCHAR(255), 
    address VARCHAR(255),
    phone VARCHAR(255),
    email VARCHAR(255)
);

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
    FOREIGN KEY (category_id) REFERENCES category(id)
);

-- Table delivery
CREATE TABLE IF NOT EXISTS delivery (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255),
    weight FLOAT,
    cost FLOAT
);

-- Table quotation
CREATE TABLE IF NOT EXISTS quotation (
    id SERIAL PRIMARY KEY,
    seller_id INT,
    customer_id INT,
    created_at TIMESTAMP DEFAULT NOW(),
    total_price FLOAT,
    FOREIGN KEY (seller_id) REFERENCES seller(id),
    FOREIGN KEY (customer_id) REFERENCES customer(id)
);

-- Table quote_product
CREATE TABLE IF NOT EXISTS quote_product (
    id SERIAL PRIMARY KEY,
    quotation_id INT,
    product_id INT,
    quantity FLOAT,
    delivery_id INT,
    FOREIGN KEY (quotation_id) REFERENCES quotation(id),
    FOREIGN KEY (product_id) REFERENCES product(id),
    FOREIGN KEY (delivery_id) REFERENCES delivery(id)
);
