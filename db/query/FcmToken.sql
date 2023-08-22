-- name: CreateFcmToken :one
INSERT INTO "FcmToken" (user_id,fcm_token,is_valid) VALUES ($1, $2, $3) RETURNING *;

-- name: GetFcmTokenById :one
SELECT * FROM "FcmToken" WHERE id = $1;

-- name: GetFcmTokenUserId :one 
SELECT * FROM "FcmToken" WHERE user_id = $1 and is_valid = 1;


-- name: DeleteFcmToken :one
DELETE FROM "FcmToken" WHERE
id=$1 RETURNING *;

-- name: UpdateFcmToken :one
UPDATE "FcmToken" 
SET fcm_token = $2, 
is_valid=$3
WHERE id = $1
RETURNING * ;


