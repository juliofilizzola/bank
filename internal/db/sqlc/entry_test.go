package db

import (
	"context"
	"testing"

	"github.com/juliofilizzola/bank/util"
	"github.com/stretchr/testify/require"
)

func TestQueries_CreateEntry(t *testing.T) {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	result, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestQueries_GetEntry(t *testing.T) {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	result, err := testQueries.CreateEntry(context.Background(), arg)

	id, err := result.LastInsertId()
	require.NoError(t, err)

	var convertId = int(id)
	n := int32(convertId)
	entry, err := testQueries.GetEntry(context.Background(), n)

	require.NoError(t, err)

	require.NotEmpty(t, entry)

	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.AccountID, entry.AccountID)

}

func TestQueries_ListEntries(t *testing.T) {
	accounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{Limit: 5,
		Offset: 1})

	require.NoError(t, err)
	require.NotEmpty(t, accounts)
}
