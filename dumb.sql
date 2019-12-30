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
    customerId INTEGER REFERENCES customers(id),
    productId INTEGER REFERENCES products(id)
);