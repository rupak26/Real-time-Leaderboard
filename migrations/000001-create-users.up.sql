-- +migrate Up
CREATE TABLE IF NOT EXISTS  users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT CHECK (age >= 0),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    occupation VARCHAR(100)
);