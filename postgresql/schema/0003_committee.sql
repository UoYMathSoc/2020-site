CREATE TABLE committee (
    id          SERIAL PRIMARY KEY,
    name        text,
    alias       varchar(255) NOT NULL,
    ordering    smallint,
    description text,
    executive   bool DEFAULT false
);
