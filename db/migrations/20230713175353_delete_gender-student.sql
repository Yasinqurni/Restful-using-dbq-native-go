-- migrate:up
ALTER TABLE students DROP COLUMN gender;

-- migrate:down
ALTER TABLE students ADD COLUMN  ENUM('L', 'P') NOT NULL;
