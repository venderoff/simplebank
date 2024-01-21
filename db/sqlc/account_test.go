package db

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	arg := CreateAccountParams{
		Owner:    "Andy",
		Balance:  121,
		Currency: "GBP",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoErrorf(t, err, "Error Identified")
	require.Equalf(t, arg.Owner, account.Owner, "Account Owner Does not Matched")
	require.Equalf(t, arg.Balance, account.Balance, "Account Balance Does Not Match")
	require.Equalf(t, arg.Currency, account.Currency, "Account Currency Does Not Match")

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
