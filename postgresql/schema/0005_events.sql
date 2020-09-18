CREATE TABLE events (
    id         SERIAL PRIMARY KEY,
    name       text NOT NULL,
    key        varchar(255) NOT NULL,
    start_time time without time zone NOT NULL,
    end_time   time without time zone,
    location   text
);
