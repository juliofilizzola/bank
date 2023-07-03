package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueries_CreateTransfers(t *testing.T) {
	arg := CreateTransfersParams{
		FromAccountID: createRandomAccount(t).ID,
		ToAccountID:   createRandomAccount(t).ID,
		Amount:        100,
	}
	result, err := testQueries.CreateTransfers(context.Background(), arg)
	// require.NoError(t, err)
	if err != nil {
		t.Errorf("err: %e", err)
	}
	require.NotEmpty(t, result)

	resultId, err := result.LastInsertId()

	require.NoError(t, err)
	require.NotEmpty(t, resultId)

	var convertIdInt = int(resultId)
	var id = int32(convertIdInt)

	res, err := testQueries.GetTransfer(context.Background(), id)

	require.NoError(t, err)

	require.NotEmpty(t, res)

	require.Equal(t, arg.Amount, res.Amount)
	require.Equal(t, arg.FromAccountID, res.FromAccountID)
	require.Equal(t, arg.ToAccountID, res.ToAccountID)
	require.NotEmpty(t, res.CreatedAt)
	require.NotEmpty(t, res.ID)

}

func TestQueries_SelectLastIntroIdTransfer(t *testing.T) {
	arg := CreateTransfersParams{
		FromAccountID: createRandomAccount(t).ID,
		ToAccountID:   createRandomAccount(t).ID,
		Amount:        100,
	}
	result, err := testQueries.CreateTransfers(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, result)

	res, err := testQueries.SelectLastIntroIdTransfer(context.Background())

	require.NoError(t, err)

	require.NotEmpty(t, res)
}

func TestQueries_GetTransfer(t *testing.T) {
	arg := CreateTransfersParams{
		FromAccountID: createRandomAccount(t).ID,
		ToAccountID:   createRandomAccount(t).ID,
		Amount:        100,
	}
	result, err := testQueries.CreateTransfers(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, result)

	resultId, err := result.LastInsertId()

	require.NoError(t, err)
	require.NotEmpty(t, resultId)

	var convertIdInt = int(resultId)
	var id = int32(convertIdInt)

	res, err := testQueries.GetTransfer(context.Background(), id)

	require.NoError(t, err)

	require.NotEmpty(t, res)

	require.Equal(t, arg.Amount, res.Amount)
	require.Equal(t, arg.FromAccountID, res.FromAccountID)
	require.Equal(t, arg.ToAccountID, res.ToAccountID)
	require.NotEmpty(t, res.CreatedAt)
	require.NotEmpty(t, res.ID)

}

func TestQueries_ListTransfersFrom(t *testing.T) {
	arg := CreateTransfersParams{
		FromAccountID: createRandomAccount(t).ID,
		ToAccountID:   createRandomAccount(t).ID,
		Amount:        100,
	}
	result, err := testQueries.CreateTransfers(context.Background(), arg)
	if err != nil {
		t.Errorf("err: %e", err)
	}

	require.NotEmpty(t, result)

	var convertIdInt = int(arg.FromAccountID)
	var id = int32(convertIdInt)

	res, err := testQueries.ListTransfersFrom(context.Background(), ListTransfersFromParams{
		FromAccountID: id,
		Limit:         4,
		Offset:        0,
	})

	require.NoError(t, err)

	require.NotEmpty(t, res)

}

func TestQueries_ListTransfersTo(t *testing.T) {
	arg := CreateTransfersParams{
		FromAccountID: createRandomAccount(t).ID,
		ToAccountID:   createRandomAccount(t).ID,
		Amount:        100,
	}
	result, err := testQueries.CreateTransfers(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, result)

	var convertIdInt = int(arg.ToAccountID)
	var id = int32(convertIdInt)

	res, err := testQueries.ListTransfersTo(context.Background(), ListTransfersToParams{
		ToAccountID: id,
		Limit:       4,
		Offset:      0,
	})

	require.NoError(t, err)

	require.NotEmpty(t, res)
}
