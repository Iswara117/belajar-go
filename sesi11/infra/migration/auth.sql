CREATE TABLE auth (
    id SERIAL PRIMARY KEY,
    email varchar(100) NOT NULL,
    password varchar(255) NOT NULL,
    craeted_at timestamptz DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW()
);