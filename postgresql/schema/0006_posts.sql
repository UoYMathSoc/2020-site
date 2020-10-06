CREATE TABLE posts (
    id       SERIAL PRIMARY KEY,
    date     timestamp with time zone DEFAULT now() NOT NULL,
    title    text NOT NULL,
    body     text,
    event_id SERIAL NOT NULL REFERENCES events(id),
    link     text
);
