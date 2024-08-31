package db

import (
	"context"
	"github.com/Dennisblay/ordering-app-server/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomUser(t *testing.T) User {
	hashPassword, err := util.HashPassword(util.RandomPassword(12))
	arg := CreateUserParams{
		FirstName:    util.RandomUser(),
		LastName:     util.RandomUser(),
		Email:        util.RandomEmail(),
		PasswordHash: hashPassword,
		Phone:        util.RandomPhone(),
		Address:      util.RandomAddress(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Address, user.Address)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUserById(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotZero(t, user2.ID)

	require.Equal(t, user2.ID, user1.ID)
	require.Equal(t, user2.FirstName, user1.FirstName)
	require.Equal(t, user2.LastName, user1.LastName)
	require.Equal(t, user2.Email, user1.Email)
	require.Equal(t, user2.Address, user1.Address)

	require.WithinDuration(t, user2.CreatedAt, user1.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	arg := UpdateUserParams{
		ID:        user1.ID,
		FirstName: util.RandomUser(),
		LastName:  util.RandomUser(),
		Email:     util.RandomEmail(),
		Phone:     util.RandomPhone(),
		Address:   util.RandomAddress(),
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user2.ID, arg.ID)
	require.Equal(t, user2.FirstName, arg.FirstName)
	require.Equal(t, user2.LastName, arg.LastName)
	require.Equal(t, user2.Email, arg.Email)
	require.Equal(t, user2.Address, arg.Address)

	require.WithinDuration(t, user2.CreatedAt, user1.CreatedAt, time.Second)
	require.NotZero(t, user2.UpdatedAt)
	require.True(t, user2.UpdatedAt.After(user1.UpdatedAt))
}

func TestDeleteUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.Error(t, err)
	require.Empty(t, user2)
}

func TestGetUsers(t *testing.T) {
	for _ = range 10 {
		CreateRandomUser(t)
	}

	args := GetUsersParams{
		Limit:  5,
		Offset: 5,
	}
	users, err := testQueries.GetUsers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, u := range users {
		require.NotEmpty(t, u)
	}
}

func TestGetUserByEmailOrPassword(t *testing.T) {
	user1 := CreateRandomUser(t)
	args := GetUserByEmailOrPasswordParams{
		ID:    user1.ID,
		Email: user1.Email,
		Phone: user1.Phone,
	}
	user2, err := testQueries.GetUserByEmailOrPassword(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user2.ID, user1.ID)
	require.Equal(t, user2.Email, user1.Email)

	require.Equal(t, user2.Address, user1.Address)
	require.Equal(t, user2.Phone, user1.Phone)

	require.WithinDuration(t, user2.CreatedAt, user1.CreatedAt, time.Second)
}

func TestUpdateUserEmail(t *testing.T) {
	user1 := CreateRandomUser(t)

	args := UpdateUserEmailParams{
		ID:    user1.ID,
		Email: util.RandomEmail(),
	}
	user2, err := testQueries.UpdateUserEmail(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.NotEqual(t, user2.Email, user1.Email)
	require.Equal(t, user2.Email, args.Email)
	require.Equal(t, user2.ID, user1.ID)
	require.Equal(t, user2.Address, user1.Address)
	require.Equal(t, user2.Phone, user1.Phone)

	require.WithinDuration(t, user2.CreatedAt, user1.CreatedAt, time.Second)
	require.NotZero(t, user2.UpdatedAt)
	require.True(t, user2.UpdatedAt.After(user1.UpdatedAt))
}

func TestUpdateAddress(t *testing.T) {
	user1 := CreateRandomUser(t)

	args := UpdateUserAddressParams{
		ID:      user1.ID,
		Address: util.RandomAddress(),
	}
	user2, err := testQueries.UpdateUserAddress(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.NotEqual(t, user2.Address, user1.Address)
	require.Equal(t, user2.Address, args.Address)
	require.Equal(t, user2.ID, user1.ID)
	require.Equal(t, user2.Email, user1.Email)
	require.Equal(t, user2.Phone, user1.Phone)

	require.WithinDuration(t, user2.CreatedAt, user1.CreatedAt, time.Second)
	require.NotZero(t, user2.UpdatedAt)
	require.True(t, user2.UpdatedAt.After(user1.UpdatedAt))
}

func TestUpdatePhone(t *testing.T) {
	user1 := CreateRandomUser(t)

	args := UpdateUserPhoneParams{
		ID:    user1.ID,
		Phone: util.RandomPhone(),
	}
	user2, err := testQueries.UpdateUserPhone(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.NotEqual(t, user2.Phone, user1.Phone)
	require.Equal(t, user2.Phone, args.Phone)
	require.Equal(t, user2.Address, user1.Address)
	require.Equal(t, user2.ID, user1.ID)
	require.Equal(t, user2.Email, user1.Email)

	require.WithinDuration(t, user2.CreatedAt, user1.CreatedAt, time.Second)
	require.NotZero(t, user2.UpdatedAt)
	require.True(t, user2.UpdatedAt.After(user1.UpdatedAt))
}

func TestUpdateUserPassword(t *testing.T) {
	user1 := CreateRandomUser(t)
	args := UpdateUserPasswordParams{
		ID:           user1.ID,
		PasswordHash: util.RandomPassword(12),
	}
	user2, err := testQueries.UpdateUserPassword(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user2.Address, user1.Address)
	require.Equal(t, user2.ID, user1.ID)
	require.Equal(t, user2.Email, user1.Email)
	require.Equal(t, user2.Phone, user1.Phone)
	require.NotEqual(t, user2.PasswordHash, user1.PasswordHash)
	require.Equal(t, user2.PasswordHash, args.PasswordHash)

	require.WithinDuration(t, user2.CreatedAt, user1.CreatedAt, time.Second)
	require.NotZero(t, user2.UpdatedAt)
	require.True(t, user2.UpdatedAt.After(user1.UpdatedAt))
}
