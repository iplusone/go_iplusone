CREATE TABLE IF NOT EXISTS messages (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  message VARCHAR(255) NOT NULL
);

INSERT INTO messages (title, message)
VALUES ('Docker E2E', 'Hello from MySQL through Go API')
ON DUPLICATE KEY UPDATE
  title = VALUES(title),
  message = VALUES(message);
