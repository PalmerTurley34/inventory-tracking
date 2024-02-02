-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, username, password)
VAlUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;