// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
}

var _ Querier = (*Queries)(nil)
