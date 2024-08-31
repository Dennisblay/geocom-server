// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUserByEmailAndPassword(ctx context.Context, arg GetUserByEmailAndPasswordParams) (GetUserByEmailAndPasswordRow, error)
	GetUserByEmailOrPassword(ctx context.Context, arg GetUserByEmailOrPasswordParams) (GetUserByEmailOrPasswordRow, error)
	GetUserById(ctx context.Context, id int64) (GetUserByIdRow, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUsersRow, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserAddress(ctx context.Context, arg UpdateUserAddressParams) (User, error)
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (User, error)
	UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (User, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error)
	UpdateUserPhone(ctx context.Context, arg UpdateUserPhoneParams) (User, error)
}

var _ Querier = (*Queries)(nil)
