-- name: CreateAnswer :one
INSERT INTO "Answer" (user_id, question_id, content) VALUES ($1, $2, $3) RETURNING *;

-- name: GetAnswerByID :one
SELECT * FROM "Answer" WHERE id = $1;

-- name: GetAnswerForUpdate :one
SELECT * FROM "Answer" WHERE id = $1 LIMIT 1 For No Key Update;

-- name: GetAnswersByQuestionID :many
SELECT * FROM "Answer" WHERE question_id = $1 and user_id=$4 ORDER BY created_at DESC LIMIT $2 OFFSET $3;


-- name: UpdateAnswersByQuestionID :one
Update "Answer"
Set content=$3, updated_at = now() WHERE id = $1 AND 
user_id=$2 
RETURNING *;


-- name: DeleteAnswerById :one
DELETE FROM "Answer"
WHERE id = $1 RETURNING *;

-- name: DeleteAnswerByQuestionId :exec
DELETE FROM "Answer" 
WHERE question_id=$1;

-- name: DeleteAnswerByUserId :exec
DELETE FROM "Answer"
WHERE user_id= $1;