-- name: CreatePoll :one
INSERT INTO polls(
    description,
    created_by
)
VALUES (
    $1,
    $2
) RETURNING id, description, created_by, is_active, created_at;

-- name: GetPolls :many
SELECT * FROM polls;

-- name: GetPoll :one
SELECT * FROM polls as p
JOIN public.poll_options po on p.id = po.poll_id
WHERE p.id=$1;

-- name: UpdatePoll :one
UPDATE polls
SET description=$2, is_active=$3
WHERE id = $1
RETURNING  *;

-- name: UpdatePollStatus :one
UPDATE polls
SET is_active=$2
WHERE id = $1
RETURNING *;

-- name: DeletePoll :one
DELETE FROM polls
WHERE id = $1
RETURNING  *;

