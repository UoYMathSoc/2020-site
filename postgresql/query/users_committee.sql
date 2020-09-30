-- name: GetUserPositions :many
SELECT *
FROM users_committee
WHERE user_id = $1;

-- name: GetPositionUsers :many
SELECT *
FROM users_committee
WHERE committee_id = $1;