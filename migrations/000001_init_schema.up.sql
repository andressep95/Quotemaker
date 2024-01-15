CREATE TABLE seller (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE customer (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  rut VARCHAR(255), 
  address VARCHAR(255),
  phone VARCHAR(255),
  email VARCHAR(255)
);

CREATE TABLE category (
  id SERIAL PRIMARY KEY,
  category_name VARCHAR(255) NOT NULL
);

CREATE TABLE product (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  category_id INT,
  length INT, 
  price INT,
  weight FLOAT,
  code VARCHAR(255),
  FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE delivery (
  id SERIAL PRIMARY KEY,
  address VARCHAR(255),
  weight FLOAT,
  cost INT
);

CREATE TABLE quotation (
  id SERIAL PRIMARY KEY,
  seller_id INT,
  customer_id INT,
  created_at TIMESTAMP DEFAULT NOW(),
  total_price INT,
  FOREIGN KEY (seller_id) REFERENCES seller(id),
  FOREIGN KEY (customer_id) REFERENCES customer(id)
);

CREATE TABLE quote_product (
  id SERIAL PRIMARY KEY,
  quotation_id INT,
  product_id INT,
  quantity INT,
  delivery_id INT,
  FOREIGN KEY (quotation_id) REFERENCES quotation(id),
  FOREIGN KEY (product_id) REFERENCES product(id),
  FOREIGN KEY (delivery_id) REFERENCES delivery(id)
);