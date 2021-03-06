CREATE DATABASE testapi;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    email VARCHAR,
    password VARCHAR
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR, 
    price DECIMAL,
    quantity INT,
    userId INTEGER REFERENCES users(id)
);

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    email VARCHAR
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    status INTEGER DEFAULT 3,
    customerid INTEGER REFERENCES customers(id),
    productid INTEGER REFERENCES products(id)
);

ALTER TABLE orders ADD COLUMN author INTEGER REFERENCES users(id);