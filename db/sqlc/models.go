// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Poll struct {
	ID          pgtype.UUID        `json:"id"`
	Description string             `json:"description"`
	IsActive    bool               `json:"is_active"`
	CreatedBy   pgtype.UUID        `json:"created_by"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

type PollOption struct {
	ID         pgtype.UUID        `json:"id"`
	PollID     pgtype.UUID        `json:"poll_id"`
	OptionText string             `json:"option_text"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
}

type User struct {
	ID        pgtype.UUID `json:"id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type Vote struct {
	ID        pgtype.UUID        `json:"id"`
	OptionID  pgtype.UUID        `json:"option_id"`
	UserID    pgtype.UUID        `json:"user_id"`
	VotedAt   pgtype.Timestamptz `json:"voted_at"`
	IpAddress string             `json:"ip_address"`
	UserAgent string             `json:"user_agent"`
}
