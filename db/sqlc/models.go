// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Poll struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedBy   uuid.UUID `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PollOption struct {
	ID         uuid.UUID `json:"id"`
	PollID     uuid.UUID `json:"poll_id"`
	OptionText string    `json:"option_text"`
	CreatedAt  time.Time `json:"created_at"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Vote struct {
	ID        uuid.UUID `json:"id"`
	OptionID  uuid.UUID `json:"option_id"`
	UserID    uuid.UUID `json:"user_id"`
	VotedAt   time.Time `json:"voted_at"`
	IpAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
}
