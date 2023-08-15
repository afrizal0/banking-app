package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDb)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5
	amount := int64(10)

	errors := make(chan error)
	results := make(chan TransferTxResult)

	// goroutines
	for i := 0; i < n; i++ {
		go func() {
			arg := TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			}
			transfer, err := store.TransferTx(context.Background(), arg)
			errors <- err
			results <- transfer

		}()
	}
	// check result
	for i := 0; i < n; i++ {
		err := <-errors
		result := <-results
		require.NoError(t, err)
		require.NotEmpty(t, result)

		// check transfers
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)

		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfers(context.Background(), transfer.ID)
		require.NoError(t, err)

		fromEntry := result.FromEntry

		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, amount, transfer.Amount)

		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntries(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry

		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)

		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntries(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// TODO: Check account balance
	}

}
