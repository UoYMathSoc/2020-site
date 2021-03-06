// Code generated by sqlc. DO NOT EDIT.

package database

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, username string) (int32, error)
	FindUserUsername(ctx context.Context, username string) (int32, error)
	GetEvent(ctx context.Context, id int32) (Event, error)
	GetPosition(ctx context.Context, id int32) (Position, error)
	GetPositionUsers(ctx context.Context, committeeID int32) ([]UsersCommittee, error)
	GetUser(ctx context.Context, id int32) (User, error)
	GetUserPositions(ctx context.Context, userID int32) ([]UsersCommittee, error)
	GetUsersPass(ctx context.Context, id int32) (UsersPass, error)
	ListEvents(ctx context.Context) ([]Event, error)
	ListPositions(ctx context.Context) ([]Position, error)
	ListUsers(ctx context.Context) ([]User, error)
	SetUsersPass(ctx context.Context, arg SetUsersPassParams) error
	UpdateUserName(ctx context.Context, arg UpdateUserNameParams) error
	UpdateUserUsername(ctx context.Context, arg UpdateUserUsernameParams) error
}

var _ Querier = (*Queries)(nil)
