CREATE TABLE IF NOT EXISTS animals (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50),
  age INT
);

INSERT INTO animals (name, age) VALUES('Hippo', 10);
INSERT INTO animals (name, age) VALUES('Ele', 10);