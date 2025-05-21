CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT NOT NULL UNIQUE,
                       password TEXT NOT NULL,
                       role TEXT
);

CREATE TABLE wines (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       year INTEGER NOT NULL,
                       price NUMERIC(10, 2) NOT NULL,
                       description TEXT
);

