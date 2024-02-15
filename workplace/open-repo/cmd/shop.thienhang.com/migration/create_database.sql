DROP TABLE IF EXISTS customer;

DROP TABLE IF EXISTS product;

DROP TABLE IF EXISTS store;

DROP TABLE IF EXISTS employee;

DROP TABLE IF EXISTS payment;

DROP TABLE IF EXISTS delivery;

DROP TABLE IF EXISTS date_order;

DROP TABLE IF EXISTS orders;

DROP TABLE IF EXISTS campaign;

DROP TABLE IF EXISTS strategy;

DROP TABLE IF EXISTS discount_type;

DROP TABLE IF EXISTS discount_value;

CREATE TABLE customer (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(200) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    city VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    street_name VARCHAR(200) NOT NULL,
    street_number VARCHAR(20) NOT NULL,
    market_segment VARCHAR(20) NOT NULL,
    isDelete BOOLEAN DEFAULT FALSE
);

-- sale_status: Product's business status (in stock, available, discontinued)
-- unit (unit of product kg, bundle, liter, can,....)
-- //online min quantity default value 0
CREATE TABLE product (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,
    unit VARCHAR(10) NOT NULL,
    category VARCHAR(20) NOT NULL,
    brand VARCHAR(50) NOT NULL,
    sale_status VARCHAR(50) NOT NULL,
    sales_type VARCHAR(50) NOT NULL,
    isDelete BOOLEAN DEFAULT FALSE
);

CREATE TABLE store (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    street_name VARCHAR(200) NOT NULL,
    street_number VARCHAR(20) NOT NULL,
    manager VARCHAR(200) NOT NULL,
    isDelete BOOLEAN DEFAULT FALSE
);

CREATE TABLE employee (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(200) NOT NULL,
    city VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    street_name VARCHAR(200) NOT NULL,
    street_number VARCHAR(20) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    age INT NOT NULL,
    salary_Level VARCHAR(20) NOT NULL,
    status VARCHAR(20) NOT NULL,
    education_level VARCHAR(20) NOT NULL,
    last_rating DECIMAL(4, 2) NOT NULL,
    isDelete BOOLEAN DEFAULT FALSE
);

CREATE TABLE payment (
    id SERIAL PRIMARY KEY NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    isDelete BOOLEAN DEFAULT FALSE
);

CREATE TABLE delivery (
    id SERIAL PRIMARY KEY NOT NULL,
    delivery_method VARCHAR(50) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    isDelete BOOLEAN DEFAULT FALSE
);

CREATE TABLE date_order (
    id SERIAL PRIMARY KEY NOT NULL,
    day_number_of_month INT NOT NULL,
    month_number_of_year INT NOT NULL,
    year INT NOT NULL,
    isDelete BOOLEAN DEFAULT FALSE
);

CREATE TABLE campaign (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    time_start VARCHAR(20) NOT NULL,
    time_end VARCHAR(20) NOT NULL,
    description VARCHAR(255) NOT NULL,
    status VARCHAR(20) NOT NULL,
    isDelete BOOLEAN DEFAULT FALSE
);

CREATE TABLE strategy (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    time_start VARCHAR(20) NOT NULL,
    time_end VARCHAR(20) NOT NULL,
    description VARCHAR(255) NOT NULL,
    status VARCHAR(20) NOT NULL,
    isDelete BOOLEAN DEFAULT FALSE
);

--invoice_value_condition is a the total amount of the order or the total of product in one brand
--quantity_condition is a quantity product in 1 orrder or 1 brand
--brand is a brand of product can have discount
--  product_id is a product id can have discount
CREATE TABLE discount_type (
    id SERIAL PRIMARY KEY NOT NULL,
    quality_discount INT NOT NULL,
    invoice_value_condition DECIMAL(10, 2) DEFAULT 0.0,
    quantity_condition INT DEFAULT 0,
    brand VARCHAR(50) DEFAULT '',
    product_id INT DEFAULT 0,
    description VARCHAR(255) NOT NULL,
    isDelete BOOLEAN DEFAULT FALSE
);

--display_name is a string name for customer can see, like a code 
CREATE TABLE discount_value (
    id SERIAL PRIMARY KEY NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    value DECIMAL(10, 2) NOT NULL,
    max_value DECIMAL(10, 2) DEFAULT 0.0,
    isDelete BOOLEAN DEFAULT FALSE
);

--Fact discout
CREATE TABLE discount (
    id SERIAL PRIMARY KEY NOT NULL,
    id_campaign INT NOT NULL,
    id_strategy INT NOT NULL,
    id_discount_type INT NOT NULL,
    id_discount_value INT NOT NULL,
    CONSTRAINT fk_discount_campaign FOREIGN KEY (id_campaign) REFERENCES campaign(id),
    CONSTRAINT fk_discount_strategy FOREIGN KEY (id_strategy) REFERENCES strategy(id),
    CONSTRAINT fk_discount_discount_type FOREIGN KEY (id_discount_type) REFERENCES discount_type(id),
    CONSTRAINT fk_discount_discount_value FOREIGN KEY (id_discount_value) REFERENCES discount_value(id),
    isDelete BOOLEAN DEFAULT FALSE
);

--Fact history change price product
-- Price changes eg 15k->10K => price change 5k
-- change status + or - sign, represents an increase or decrease in price
--New price is price after change
CREATE TABLE price_history (
    id SERIAL PRIMARY KEY NOT NULL,
    product_id INT NOT NULL,
    price_change DECIMAL(10, 2) NOT NULL,
    change_status VARCHAR(2) NOT NULL,
    new_price DECIMAL(10, 2) NOT NULL,
    date_change_price TIMESTAMP,
    CONSTRAINT fk_price_history_product FOREIGN KEY (product_id) REFERENCES product(id),
    isDelete BOOLEAN DEFAULT FALSE
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY NOT NULL,
    customer_id INT DEFAULT 0,
    product_id INT NOT NULL,
    store_id INT NOT NULL,
    employee_id INT NOT NULL,
    payment_id INT NOT NULL,
    delivery_id INT DEFAULT 0,
    date_order_id INT NOT NULL,
    discount_id INT DEFAULT 0,
    total_amount_order DECIMAL(12, 2) NOT NULL,
    discount_amount DECIMAL(12, 2) DEFAULT 0,
    isDelete BOOLEAN DEFAULT FALSE,
    CONSTRAINT fk_order_product FOREIGN KEY (product_id) REFERENCES product(id),
    CONSTRAINT fk_order_store FOREIGN KEY (store_id) REFERENCES store(id),
    CONSTRAINT fk_order_employee FOREIGN KEY (employee_id) REFERENCES employee(id),
    CONSTRAINT fk_order_payment FOREIGN KEY (payment_id) REFERENCES payment(id),
    CONSTRAINT fk_order_date FOREIGN KEY (date_order_id) REFERENCES date_order(id)
);