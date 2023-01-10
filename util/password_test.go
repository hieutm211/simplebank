package util

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := gofakeit.Name()

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = VerifyPassword(hashedPassword, password)
	require.NoError(t, err)

	wrongPassword := gofakeit.Name()
	err = VerifyPassword(hashedPassword, wrongPassword)
	require.Error(t, err)

	hashedPassword2, err := HashPassword(password)
	require.NotEqual(t, hashedPassword, hashedPassword2)
}
