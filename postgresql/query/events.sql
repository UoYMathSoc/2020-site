-- name: ListEvents :many
SELECT *
FROM events;

-- name: GetEvent :one
SELECT *
FROM events
WHERE id = $1;
