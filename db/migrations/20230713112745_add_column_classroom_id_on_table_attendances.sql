-- migrate:up
ALTER TABLE attendances ADD COLUMN classroom_id INT;
ALTER TABLE attendances ADD CONSTRAINT classroom_id 
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id);

-- migrate:down
ALTER TABLE attendances DROP COLUMN classroom_id INT;
