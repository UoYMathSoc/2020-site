-- name: ListPositions :many
SELECT *
FROM committee;

-- name: GetPosition :one
SELECT *
FROM committee
WHERE id = $1;