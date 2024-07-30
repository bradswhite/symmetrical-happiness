CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  username TEXT NOT NULL,
  email TEXT NOT NULL,
  encrypted_password TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(email, username)
);

CREATE TABLE IF NOT EXISTS software (
  id UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  name TEXT NOT NULL,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  image TEXT NOT NULL,
  url TEXT NOT NULL,
  user_id UUID,
  username TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(name, title)
);

CREATE TABLE IF NOT EXISTS software_likes (
  software_id UUID NOT NULL,
  user_id UUID NOT NULL,
  username text NOT NULL, 
  liked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(software_id, username)
);
