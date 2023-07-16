-- migrate:up
ALTER TABLE students ADD COLUMN gender ENUM('L', 'P') NOT NULL;

-- migrate:down
ALTER TABLE students DROP COLUMN gender;
