// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package gen

import (
	"context"
)

type Querier interface {
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Account, error)
	CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entry, error)
	CreateTransfers(ctx context.Context, arg CreateTransfersParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTransfers(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAccount(ctx context.Context, arg ListAccountParams) ([]Account, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	createTeacher(ctx context.Context, arg createTeacherParams) (Teacher, error)
	deleteAccount(ctx context.Context, id int64) error
}

var _ Querier = (*Queries)(nil)
