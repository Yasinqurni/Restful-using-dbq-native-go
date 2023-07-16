-- migrate:up
ALTER TABLE classrooms CHANGE  capacity room_size INT NOT NULL;

-- migrate:down
ALTER TABLE classrooms CHANGE room_size capacity INT NOT NULL;
