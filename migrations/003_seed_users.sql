INSERT INTO users (username, password_hash, role)
VALUES 
  ('admin', '$2b$10$6OqfD3vM6VvXukcK0Y9d6e0rP7XkM/U9IwhmJ0pNlz/6FqNwWcN2e', 'admin') 
ON DUPLICATE KEY UPDATE username = username;
-- admin123