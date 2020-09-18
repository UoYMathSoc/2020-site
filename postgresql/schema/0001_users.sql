CREATE TABLE users (
    id       SERIAL PRIMARY KEY,
    username text NOT NULL,
    name     text NOT NULL,
    bio      text
);
