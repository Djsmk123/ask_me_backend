-- name: CreateUser :one
INSERT INTO "User" (username, email, password_hash) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM "User" WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM "User" WHERE email = $1;

-- name: