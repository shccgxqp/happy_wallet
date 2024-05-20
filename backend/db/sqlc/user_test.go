package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shccgxqp/happy_wallet/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username: util.RandomUsername(),
		Password: hashedPassword,
		Email:    util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), arg)
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

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
	require.WithinDuration(t, user1.UpdatedAt.Time, user2.UpdatedAt.Time, time.Second)
}

func TestUpdateUser(t *testing.T) {
	oldUser := createRandomUser(t)

	newUsername := util.RandomUsername()
	newEmail := util.RandomEmail()
	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		ID: oldUser.ID,
		Username: pgtype.Text{
			String: newUsername,
			Valid:  true,
		},
		Password: pgtype.Text{
			String: newHashedPassword,
			Valid:  true,
		},
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Password, updatedUser.Password)
	require.Equal(t, newHashedPassword, updatedUser.Password)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.NotEqual(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, newUsername, updatedUser.Username)
}

func TestUpdateUserOnlyUsername(t *testing.T) {
	oldUser := createRandomUser(t)

	newUsername := util.RandomUsername()
	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		ID: oldUser.ID,
		Username: pgtype.Text{
			String: newUsername,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, newUsername, updatedUser.Username)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.Password, updatedUser.Password)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	oldUser := createRandomUser(t)

	newEmail := util.RandomEmail()
	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		ID: oldUser.ID,
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.Equal(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, oldUser.Password, updatedUser.Password)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)

	newPassword := util.RandomPassword()
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		ID: oldUser.ID,
		Password: pgtype.Text{
			String: newHashedPassword,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Password, updatedUser.Password)
	require.Equal(t, newHashedPassword, updatedUser.Password)
	require.Equal(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, oldUser.Email, updatedUser.Email)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testStore.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testStore.GetUser(context.Background(), user1.Username)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testStore.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
		require.NotEmpty(t, user.ID)
		require.NotEmpty(t, user.Username)
		require.NotEmpty(t, user.Password)
		require.NotEmpty(t, user.Email)
		require.NotZero(t, user.CreatedAt)
		require.NotZero(t, user.UpdatedAt)
	}
}
