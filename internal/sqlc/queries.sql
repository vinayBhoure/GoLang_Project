-- name: CreateUser :one
INSERT INTO "users"(first_name, last_name)
VALUES ($1, $2) RETURNING *;

-- name: GetUser :one
SELECT * FROM "users"
WHERE id = $1;

-- name: ListUser :many
SELECT id, first_name, last_name FROM "users";

-- name: UpdateUser :exec
UPDATE "users"
SET 
  first_name = $1,
  last_name = $2
WHERE id = $3;

-- name: DeleteUser :exec
DELETE FROM "users"
WHERE id = $1;

-- name: CreateTweet :one
INSERT INTO "tweets"(content, "user")
VALUES ($1, $2) RETURNING *;

-- name: GetTweet :one
SELECT * FROM "tweets"
WHERE id = $1;

-- name: ListTweet :many
SELECT id, content, "user" FROM "tweets";

-- name: UpdateTweet :exec
UPDATE "tweets"
SET 
  content = $1,
  "user" = $2
WHERE id = $3;

-- name: DeleteTweet :exec
DELETE FROM "tweets"
WHERE id = $1 AND user = $2;