-- name: GetAllUsers :many
SELECT * FROM users;


-- name: CreateUser :one
INSERT INTO users (username, password_hash, email)
VALUES ($1, $2, $3)
RETURNING *;

