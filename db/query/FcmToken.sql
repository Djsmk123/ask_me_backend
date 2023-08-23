-- name: CreateFcmToken :one
INSERT INTO "FcmToken" (id,user_id,fcm_token,is_valid) VALUES ($1, $2, $3,$4) RETURNING *;

-- name: GetFcmTokenById :one
SELECT * FROM "FcmToken" WHERE id = $1;

-- name: GetFcmTokenUserId :one 
SELECT * FROM "FcmToken" WHERE user_id = $1 and is_valid = 1;


-- name: DeleteFcmToken :one
DELETE FROM "FcmToken" WHERE
id=$1 RETURNING *;

-- name: DeleteFcmTokenByUserId :many
DELETE FROM "FcmToken" WHERE user_id = $1
RETURNING *;

-- name: UpdateFcmToken :one
UPDATE "FcmToken" 
SET is_valid=$2
WHERE id = $1
RETURNING * ;


