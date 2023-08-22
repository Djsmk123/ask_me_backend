-- name: CreateJwtToken :one
INSERT INTO "Token" (user_id,jwt_token,is_valid) VALUES ($1, $2, $3) RETURNING *;

-- name: GetJwtTokenById :one
SELECT * FROM "Token" WHERE id = $1;

-- name: GetJwtTokenUserId :one 
SELECT * FROM "Token" WHERE jwt_token ILIKE $2 and user_id = $1;


-- name: DeleteJwtToken :one
DELETE FROM "Token" WHERE
id=$1 RETURNING *;

-- name: UpdateJwtToken :one
UPDATE "Token" 
SET jwt_token = $2, 
is_valid=$3
WHERE id = $1
RETURNING * ;


