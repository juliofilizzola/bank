package db

import (
	"context"
	"testing"

	"bank/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func TestQueries_CreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	account, err := testQueries.SelectLastIntroID(context.Background())
	require.NoError(t, err)

	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	require.NotEmpty(t, account.CreatedAt)
}

func TestQueries_GetAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "tom",
		Balance:  100,
		Currency: "USD",
	}
	account, err := testQueries.GetAccount(context.Background(), 3)

	require.NoError(t, err)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotEmpty(t, account.CreatedAt)
}

func TestQueries_ListAccounts(t *testing.T) {
	accounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{Limit: 5,
		Offset: 1})

	require.NoError(t, err)
	require.NotEmpty(t, accounts)
}

func TestQueries_DeleteAccount(t *testing.T) {
	err := testQueries.DeleteAccount(context.Background(), 1)

	require.NoError(t, err)
}

func TestQueries_UpdatedAccounts(t *testing.T) {
	arg := CreateAccountParams{
		Balance: 200,
	}

	err := testQueries.UpdatedAccounts(context.Background(), UpdatedAccountsParams{
		ID: 2, Balance: arg.Balance,
	})

	require.NoError(t, err)

	accounts, err := testQueries.GetAccount(context.Background(), 2)

	require.NoError(t, err)

	require.Equal(t, arg.Balance, accounts.Balance)
}
