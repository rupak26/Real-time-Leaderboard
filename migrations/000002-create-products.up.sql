-- +migrate Up
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    productname VARCHAR(100),
	url TEXT,
	quantity INT
);