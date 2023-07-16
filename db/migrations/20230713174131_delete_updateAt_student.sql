-- migrate:up
ALTER TABLE students DROP COLUMN updated_at;

-- migrate:down
ALTER TABLE students ADD COLUMN  updated_at TIMESTAMP NULL;
