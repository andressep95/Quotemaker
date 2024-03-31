CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Table category
CREATE TABLE IF NOT EXISTS category (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    category_name VARCHAR(255) NOT NULL
);

-- Table product
CREATE TABLE IF NOT EXISTS product (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    category_id UUID,
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
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
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
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    quotation_id UUID,
    product_id UUID,
    quantity FLOAT,
    FOREIGN KEY (quotation_id) REFERENCES quotation(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product(id)
);
-- Agrega una restricción única en la tabla quote_product
ALTER TABLE quote_product
ADD CONSTRAINT uk_quote_product_quotation_product UNIQUE (quotation_id, product_id);



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
