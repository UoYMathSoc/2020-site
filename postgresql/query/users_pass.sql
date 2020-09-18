-- name: GetUsersPass :one
SELECT *
FROM users_pass
WHERE id = $1;

-- name: SetUsersPass :exec
INSERT INTO users_pass (id, password)
VALUES ($1, $2);
