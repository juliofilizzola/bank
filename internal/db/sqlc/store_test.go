package db

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func TestStore_TransferTx(t *testing.T) {

	store := NewStore(testDb)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 1
	const amount int64 = 10

	errs := make(chan error)
	results := make(chan TransferTxResult)
	account, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	var realValueBalance = account.Balance - amount
	for i := 0; i < n; i++ {

		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountId:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result

		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs

		require.NoError(t, err)

		results := <-results
		require.NotEmpty(t, results)

		transfer := results.Transfer

		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
		fmt.Println(transfer.ID)
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		fromEntry := results.FromEntry

		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

	}
	account4, err := store.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.Equal(t, account4.Balance, realValueBalance)
}
