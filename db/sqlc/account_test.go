package db

import (
	"context"
	"testing"

	"github.com/MDelarosa1993/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	params := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, params.Owner, account.Owner)
	require.Equal(t, params.Balance, account.Balance)
	require.Equal(t, params.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
