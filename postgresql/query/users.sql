-- name: ListUsers :many
SELECT *
FROM users;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (username)
VALUES ($1)
RETURNING id;

-- name: UpdateUserName :exec
UPDATE users
SET name = $2
WHERE id = $1;

-- name: UpdateUserUsername :exec
UPDATE users
SET username = $2
WHERE id = $1;

-- name: FindUserUsername :one
SELECT id
FROM users
WHERE username LIKE $1;
