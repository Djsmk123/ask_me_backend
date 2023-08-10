-- name: CreateAnswer :one
INSERT INTO "Answer" (user_id, question_id, content) VALUES ($1, $2, $3) RETURNING *;

-- name: GetAnswerByID :one
SELECT * FROM "Answer" WHERE id = $1;

-- name: GetAnswerForUpdate :one
SELECT * FROM "Answer" WHERE id = $1 LIMIT 1 For No Key Update;

-- name: GetAnswersByQuestionID :many
SELECT * FROM "Answer" WHERE question_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;


-- name: UpdateAnswersByQuestionID :one
Update "Answer"
Set content=$2 WHERE id = $1 
RETURNING *;


-- name: AnswerDelete :exec
DELETE FROM "Answer"
WHERE id = $1;

-- name: DeleteAnswerByQuestionId :exec

DELETE FROM "Answer"
WHERE question_id=$1;