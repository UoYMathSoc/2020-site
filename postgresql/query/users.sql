-- name: ListUsers :many
SELECT *
FROM users;

-- name: Get :one
SELECT *
FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (username, name)
VALUES ($1, $1)
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
