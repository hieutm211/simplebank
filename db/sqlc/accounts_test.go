package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    gofakeit.Name(),
		Balance:  int64(gofakeit.Number(1000, 10000)),
		Currency: gofakeit.CurrencyShort(),
	}

	gotAccount, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, gotAccount)

	require.NotZero(t, gotAccount.ID)
	require.Equal(t, arg.Owner, gotAccount.Owner)
	require.Equal(t, arg.Balance, gotAccount.Balance)
	require.Equal(t, arg.Currency, gotAccount.Currency)
	require.NotZero(t, gotAccount.CreatedAt)

	return gotAccount
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	gotAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotAccount)

	require.Equal(t, createdAccount.ID, gotAccount.ID)
	require.Equal(t, createdAccount.Owner, gotAccount.Owner)
	require.Equal(t, createdAccount.Balance, gotAccount.Balance)
	require.Equal(t, createdAccount.Currency, gotAccount.Currency)
	require.NotZero(t, createdAccount.CreatedAt, gotAccount.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      createdAccount.ID,
		Balance: int64(gofakeit.Number(0, 1000)),
	}

	gotAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, gotAccount)

	require.Equal(t, createdAccount.ID, gotAccount.ID)
	require.Equal(t, createdAccount.Owner, gotAccount.Owner)
	require.Equal(t, arg.Balance, gotAccount.Balance)
	require.Equal(t, createdAccount.Currency, gotAccount.Currency)
	require.NotZero(t, createdAccount.CreatedAt, gotAccount.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)

	gotAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, gotAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	listAccounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{Limit: 5, Offset: 5})
	require.NoError(t, err)
	require.NotEmpty(t, listAccounts)
	require.Len(t, listAccounts, 5)
}
