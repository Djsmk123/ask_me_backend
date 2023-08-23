-- name: CreateJwtToken :one
INSERT INTO "Token" (user_id,jwt_token,created_at,expires_at) VALUES ($1, $2, $3,$4) RETURNING *;

-- name: GetJwtTokenById :one
SELECT * FROM "Token" WHERE id = $1;

-- name: GetJwtTokenUserId :one 
SELECT * FROM "Token" WHERE jwt_token ILIKE $2 and user_id = $1;


-- name: DeleteJwtToken :one
DELETE FROM "Token" WHERE
id=$1 RETURNING *;


-- name: DeleteJWTokenByUserId :many
DELETE FROM "Token" WHERE user_id = $1
RETURNING *;

-- name: UpdateJwtToken :one
UPDATE "Token" 
SET expires_at=now()
WHERE jwt_token ILIKE $1 and user_id = $2
RETURNING * ;


