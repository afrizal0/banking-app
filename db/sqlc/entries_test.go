package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntries(t *testing.T) Entry {
	account1 := createRandomAccount(t)
	arg := CreateEntriesParams{
		AccountID: account1.ID,
		Amount:    10000,
	}
	entry, err := testQueries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	return entry
}

func TestCreateEntries(t *testing.T) {
	createRandomEntries(t)
}

func TestGetEntries(t *testing.T) {
	entries1 := createRandomEntries(t)
	entries2, err := testQueries.GetEntries(context.Background(), entries1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entries2)

	require.Equal(t, entries1.ID, entries2.ID)
	require.Equal(t, entries1.AccountID, entries2.AccountID)
	require.Equal(t, entries1.Amount, entries2.Amount)

	require.WithinDuration(t, entries1.CreatedAt, entries2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntries(t)
	}

	arg := GetListEntriesParams{
		Limit:  5,
		Offset: 0,
	}

	entries, err := testQueries.GetListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}
