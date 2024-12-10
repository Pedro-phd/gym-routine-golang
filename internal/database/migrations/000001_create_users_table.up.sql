CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  supabase_id VARCHAR(255) NOT NULL UNIQUE
);

CREATE UNIQUE INDEX idx_supabase_id ON users(supabase_id);