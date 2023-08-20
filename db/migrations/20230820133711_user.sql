-- migrate:up
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR UNSIGNED NOT NULL PRIMARY KEY,
    name CHAR(36) NOT NULL,
    email CHAR(36) NOT NULL,
    phone CHAR(36) NOT NULL,
    password CHAR(36) NOT NULL,
    role CHAR(36) NOT NULL,
    gender CHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL ,
    deleted_at TIMESTAMP NULL , PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE users;
