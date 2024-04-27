-- name: GetSubscriber :one
SELECT *
FROM subscribers
WHERE email = $1;

-- name: UpsertSubscriber :one
INSERT INTO subscribers (email, active)
VALUES ($1, $2) ON CONFLICT(email)
DO
UPDATE SET active = $2, updated_at = NOW()
    returning *;

-- name: GetActiveSubscribersWithPaginate :many
SELECT *
FROM subscribers
WHERE active = true
ORDER BY created_at ASC LIMIT $1
OFFSET $2;

-- name: CountActiveSubscriber :one
SELECT COUNT(*)
FROM subscribers
WHERE active = true;