-- name: CreateTweet :one
INSERT INTO "tweets"(content, userid)
VALUES ($1, $2) RETURNING *;

-- name: GetTweet :one
SELECT * FROM "tweets"
WHERE id = $1;

-- name: ListTweet :many
SELECT id, content FROM "tweets" WHERE userid = $1;

-- name: UpdateTweet :exec
UPDATE "tweets"
SET 
  content = $1
WHERE id = $2;

-- name: DeleteTweet :exec
DELETE FROM "tweets"
WHERE id = $1;