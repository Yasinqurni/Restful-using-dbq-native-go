-- migrate:up
CREATE TABLE IF NOT EXISTS items (
    id VARCHAR UNSIGNED NOT NULL PRIMARY KEY,
    user_id VARCHAR UNSIGNED NOT NULL,
    category CHAR(36) NOT NULL,
    price DOUBLE NOT NULL,
    quantity DOUBLE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL, 
    FOREIGN KEY (user_id) REFERENCES user(id)
)


-- migrate:down
DROP TABLE items;
