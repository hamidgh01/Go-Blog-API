CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(64) NOT NULL UNIQUE,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    bio VARCHAR(500),
    enabled BOOLEAN DEFAULT true,
    isSuperuser BOOLEAN DEFAULT false,
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    modifiedAt TIMESTAMP WITH TIME ZONE
);

-- NOTE: `UNIQUE` keyword in table definition, automatically creates index (check at psql)
-- CREATE UNIQUE INDEX idx_users_username ON users (username);
-- CREATE UNIQUE INDEX idx_users_email ON users (email);
