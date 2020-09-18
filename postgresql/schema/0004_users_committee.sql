CREATE TABLE users_committee (
    id           SERIAL PRIMARY KEY,
    user_id      SERIAL NOT NULL REFERENCES users(id),
    committee_id SERIAL NOT NULL REFERENCES committee(id),
    from_date    date NOT NULL,
    till_date    date
);
