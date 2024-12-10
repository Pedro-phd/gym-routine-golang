-- name: CreateUser :exec
INSERT INTO users (name, supabase_id) VALUES ($1, $2);