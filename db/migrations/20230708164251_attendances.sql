-- migrate:up
CREATE TABLE IF NOT EXISTS attendances (
  id INT PRIMARY KEY AUTO_INCREMENT,
  date DATETIME NOT NULL,
  attendance BOOLEAN NOT NULL,
  student_id INT NOT NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL ,
  deleted_at TIMESTAMP NULL,
  
  FOREIGN KEY (student_id) REFERENCES students(id)
);

-- migrate:down
DROP TABLE attendances;
