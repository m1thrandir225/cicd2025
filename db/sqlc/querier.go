// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreatePoll(ctx context.Context, arg CreatePollParams) (CreatePollRow, error)
	CreatePollOption(ctx context.Context, arg CreatePollOptionParams) (PollOption, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	CreateVote(ctx context.Context, arg CreateVoteParams) (Vote, error)
	DeletePollOption(ctx context.Context, id uuid.UUID) (PollOption, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (DeleteUserRow, error)
	DeleteVote(ctx context.Context, id uuid.UUID) (Vote, error)
	DisablePoll(ctx context.Context, arg DisablePollParams) (Poll, error)
	GetUserDetails(ctx context.Context, id uuid.UUID) (User, error)
	UpdateVoteOption(ctx context.Context, arg UpdateVoteOptionParams) (Vote, error)
}

var _ Querier = (*Queries)(nil)
