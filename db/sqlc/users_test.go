package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       gofakeit.Username(),
		FullName:       gofakeit.Name(),
		Email:          gofakeit.Email(),
		HashedPassword: "secret",
	}

	gotUser, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, gotUser)

	require.NotZero(t, gotUser.Username)
	require.Equal(t, arg.FullName, gotUser.FullName)
	require.Equal(t, arg.Email, gotUser.Email)
	require.Equal(t, arg.HashedPassword, gotUser.HashedPassword)
	require.NotZero(t, gotUser.CreatedAt)
	require.True(t, gotUser.PasswordChangedAt.IsZero())

	return gotUser
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	createdUser := createRandomUser(t)

	gotUser, err := testQueries.GetUser(context.Background(), createdUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, gotUser)

	require.Equal(t, createdUser.Username, gotUser.Username)
	require.Equal(t, createdUser.FullName, gotUser.FullName)
	require.Equal(t, createdUser.Email, gotUser.Email)
	require.Equal(t, createdUser.HashedPassword, gotUser.HashedPassword)
	require.NotZero(t, createdUser.CreatedAt, gotUser.CreatedAt)
	require.True(t, createdUser.PasswordChangedAt.IsZero())
}
