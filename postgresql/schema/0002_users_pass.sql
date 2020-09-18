CREATE TABLE users_pass (
    id       SERIAL PRIMARY KEY REFERENCES users(id),
    password VARCHAR(255) NOT NULL
);
