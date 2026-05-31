CREATE TABLE IF NOT EXISTS accounts (
  id      INT AUTO_INCREMENT PRIMARY KEY,
  name    VARCHAR(100) NOT NULL,
  balance INT          NOT NULL DEFAULT 0
);

INSERT INTO accounts (name, balance) VALUES
  ('Alice', 1000),
  ('Bob',   500);
