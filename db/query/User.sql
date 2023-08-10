-- name: CreateUser :one
INSERT INTO "User" (username, email, password_hash,public_profile_image,private_profile_image,provider) VALUES ($1, $2, $3,$4,$5,$6) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM "User" WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM "User" WHERE email = $1;

-- name: DeleteUserById :one
DELETE FROM "User" WHERE
id=$1 RETURNING *;

-- name: UpdateUserPassword :one
UPDATE "User"
SET password_hash = $2, updated_at = now()
WHERE id = $1
RETURNING *;
