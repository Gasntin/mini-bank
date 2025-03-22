package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Gasntin/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	getAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getAccount)

	require.Equal(t, account.ID, getAccount.ID)
	require.Equal(t, account.Owner, getAccount.Owner)
	require.Equal(t, account.Balance, getAccount.Balance)
	require.Equal(t, account.Currency, getAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, getAccount.CreatedAt, time.Second)
}

func TestUpdateAccountBalance(t *testing.T) {
	accout := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      accout.ID,
		Balance: util.RandomMoney(),
	}
	resultUpdate, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, resultUpdate)

	require.Equal(t, accout.ID, resultUpdate.ID)
	require.Equal(t, arg.Balance, resultUpdate.Balance)
	require.Equal(t, accout.Owner, resultUpdate.Owner)
	require.Equal(t, accout.Currency, resultUpdate.Currency)
	require.WithinDuration(t, accout.CreatedAt, resultUpdate.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	getAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, getAccount)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestListAccount(t *testing.T) {
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
