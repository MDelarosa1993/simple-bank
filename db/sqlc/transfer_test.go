package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/MDelarosa1993/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	params := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:    account2.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, params.FromAccountID, transfer.FromAccountID)
	require.Equal(t, params.ToAccountID, transfer.ToAccountID)
	require.Equal(t, params.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)

	deletedTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedTransfer)
}

func TestListTransfers(t *testing.T) {
	for range 10 {
		createRandomTransfer(t)
	}

	params := ListTransfersParams{
		Limit: 5,
		Offset: 5,
	}

	entries, err := testQueries.ListTransfers(context.Background(), params)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestUpdateTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	params := UpdateTransferParams{
		ID:      transfer.ID,
		Amount: util.RandomMoney(),
	}

	updatedTransfer, err := testQueries.UpdateTransfer(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, updatedTransfer)

	require.Equal(t, transfer.ID, updatedTransfer.ID)
	require.Equal(t, transfer.FromAccountID, updatedTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, updatedTransfer.ToAccountID)
	require.Equal(t, params.Amount, updatedTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, updatedTransfer.CreatedAt, time.Second)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}