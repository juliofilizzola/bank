package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreateEntry(t *testing.T) {
	arg := CreateEntryParams{
		AccountID: 2,
		Amount:    320,
	}

	err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
}

func TestQueries_GetEntry(t *testing.T) {
	arg := CreateEntryParams{
		AccountID: 2,
		Amount:    320,
	}

	entry, err := testQueries.GetEntry(context.Background(), 1)

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
