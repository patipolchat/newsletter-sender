-- name: CreateSendingEmails :batchmany
INSERT INTO sending_emails (subscriber_id, newsletter_id, status)
    VALUES ($1, $2, 'pending')
    RETURNING *, (SELECT email FROM subscribers WHERE id = $1) AS email, (SELECT header FROM newsletters WHERE id = $2) AS header, (SELECT body FROM newsletters WHERE id = $2) AS body;

-- name: UpdateSendingEmailsStatus :one
UPDATE sending_emails
SET status     = $2,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: BatchUpdateSendingEmailsStatus :batchmany
UPDATE sending_emails
SET status     = $2,
    updated_at = now()
WHERE id = $1
RETURNING *;