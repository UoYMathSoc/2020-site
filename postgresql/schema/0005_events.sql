CREATE TABLE events (
    id          SERIAL PRIMARY KEY,
    name        text NOT NULL,
    start_time  TIMESTAMPTZ NOT NULL,
    end_time    TIMESTAMPTZ,
    location    text,
    description text
);
