// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: poll_options.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPollOption = `-- name: CreatePollOption :one
INSERT INTO poll_options(
    poll_id,
    option_text
) VALUES (
    $1::uuid,
    $2::text
) RETURNING id, poll_id, option_text, created_at
`

type CreatePollOptionParams struct {
	PollID     uuid.UUID `json:"poll_id"`
	OptionText string    `json:"option_text"`
}

func (q *Queries) CreatePollOption(ctx context.Context, arg CreatePollOptionParams) (PollOption, error) {
	row := q.db.QueryRow(ctx, createPollOption, arg.PollID, arg.OptionText)
	var i PollOption
	err := row.Scan(
		&i.ID,
		&i.PollID,
		&i.OptionText,
		&i.CreatedAt,
	)
	return i, err
}

const deletePollOption = `-- name: DeletePollOption :one
DELETE from poll_options
WHERE id = $1::uuid
RETURNING id, poll_id, option_text, created_at
`

func (q *Queries) DeletePollOption(ctx context.Context, id uuid.UUID) (PollOption, error) {
	row := q.db.QueryRow(ctx, deletePollOption, id)
	var i PollOption
	err := row.Scan(
		&i.ID,
		&i.PollID,
		&i.OptionText,
		&i.CreatedAt,
	)
	return i, err
}

const getOption = `-- name: GetOption :one
SELECT id, poll_id, option_text, created_at FROM poll_options
WHERE id=$1
LIMIT 1
`

func (q *Queries) GetOption(ctx context.Context, id uuid.UUID) (PollOption, error) {
	row := q.db.QueryRow(ctx, getOption, id)
	var i PollOption
	err := row.Scan(
		&i.ID,
		&i.PollID,
		&i.OptionText,
		&i.CreatedAt,
	)
	return i, err
}

const updatePollOption = `-- name: UpdatePollOption :one
UPDATE poll_options
SET option_text=$2
WHERE id=$1
RETURNING id, poll_id, option_text, created_at
`

type UpdatePollOptionParams struct {
	ID         uuid.UUID `json:"id"`
	OptionText string    `json:"option_text"`
}

func (q *Queries) UpdatePollOption(ctx context.Context, arg UpdatePollOptionParams) (PollOption, error) {
	row := q.db.QueryRow(ctx, updatePollOption, arg.ID, arg.OptionText)
	var i PollOption
	err := row.Scan(
		&i.ID,
		&i.PollID,
		&i.OptionText,
		&i.CreatedAt,
	)
	return i, err
}
