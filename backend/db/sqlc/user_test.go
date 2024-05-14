package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/shccgxqp/happt_wallet/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username: util.RandomUsername(),
		Password: hashedPassword,
		Email: util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(),arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetAccount(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt,time.Second)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt,time.Second)
}

func TestUpdateAccount(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:       user1.ID,
		Username: util.RandomUsername(),
		Password: util.RandomPassword(),
		Email:    util.RandomEmail(),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t,user2 )

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Password, user2.Password)
	require.Equal(t, arg.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt,time.Second)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt,time.Second)
}

func TestDeleteAccount(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.Error(t, err)
	require.EqualError(t,err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListAccounts(t *testing.T) {
	for i :=0; i < 10; i++{
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit: 5,
		Offset: 5,
	}
	
	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _,user := range users{
		require.NotEmpty(t, user)
		require.NotEmpty(t, user.ID)
		require.NotEmpty(t, user.Username)
		require.NotEmpty(t, user.Password)
		require.NotEmpty(t, user.Email)
		require.NotZero(t, user.CreatedAt)
		require.NotZero(t, user.UpdatedAt)
	}
}