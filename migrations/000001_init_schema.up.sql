CREATE TABLE seller (
  id INT PRIMARY KEY,
  name VARCHAR
);

CREATE TABLE customer (
  id INT PRIMARY KEY,
  name VARCHAR,
  rut VARCHAR, 
  address VARCHAR,
  phone VARCHAR,
  email VARCHAR
);

CREATE TABLE category (
  id INT PRIMARY KEY,
  category_name VARCHAR
);

CREATE TABLE product (
  id INT PRIMARY KEY,
  name VARCHAR,
  category_id INT,
  length VARCHAR, 
  price INT,
  weight INT,
  code VARCHAR,
  FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE delivery (
  id INT PRIMARY KEY,
  address VARCHAR,
  weight INT,
  cost INT
);

CREATE TABLE quotation (
  id INT PRIMARY KEY,
  seller_id INT,
  customer_id INT,
  date TIMESTAMP DEFAULT (NOW()),
  total_price INT,
  FOREIGN KEY (seller_id) REFERENCES seller(id),
  FOREIGN KEY (customer_id) REFERENCES customer(id)
);

CREATE TABLE quote_product (
  id INT PRIMARY KEY,
  quotation_id INT,
  product_id INT,
  quantity INT,
  delivery_id INT,
  FOREIGN KEY (quotation_id) REFERENCES quotation(id),
  FOREIGN KEY (product_id) REFERENCES product(id),
  FOREIGN KEY (delivery_id) REFERENCES delivery(id)
);


