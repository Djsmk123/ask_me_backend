-- name: CreateQuestion :one
INSERT INTO "Question" (user_id, content) VALUES ($1, $2) RETURNING *;

-- name: GetQuestionByID :one
SELECT * FROM "Question" WHERE id = $1;



-- name: GetQuestionForUpdate :one
SELECT * FROM "Question" WHERE id = $1 LIMIT 1 For No Key Update;

-- name: GetQuestionsByUserID :many
SELECT *
FROM "Question"
WHERE "user_id" = $1
AND ("content" ILike sqlc.narg('content') OR sqlc.narg('content') IS NULL)
ORDER BY "created_at" DESC
LIMIT $2
OFFSET $3;

-- name: UpdateQuestionById :one
Update "Question"
Set content=$3, updated_at = now()
WHERE id = $1 and user_id= $2

RETURNING *;


-- name: QuestionDelete :one
DELETE FROM "Question"
WHERE id = $1 
RETURNING *;

-- name: DeleteQuestionByUserId :exec
DELETE FROM "Question"
WHERE user_id= $1;

