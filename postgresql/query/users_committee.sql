-- name: GetUserPositions :many
SELECT committee.*
FROM committee
    INNER JOIN users_committee uc on committee.id = uc.committee_id
    INNER JOIN users u on uc.user_id = u.id
WHERE u.id = $1;

-- name: GetPositionUsers :many
SELECT users.*
FROM users
    INNER JOIN users_committee uc on users.id = uc.user_id
    INNER JOIN committee c on uc.committee_id = c.id
WHERE c.id = $1;