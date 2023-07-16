-- migrate:up
ALTER TABLE students DROP COLUMN age;

-- migrate:down
ALTER TABLE students ADD COLUMN age INT NOT NULL;
