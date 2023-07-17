-- migrate:up
CREATE TABLE IF NOT EXISTS students (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(26) NOT NULL,
  age INT NOT NULL,
  date_of_birth DATE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL ,
  deleted_at TIMESTAMP NULL 
);

-- migrate:down
DROP TABLE students;
