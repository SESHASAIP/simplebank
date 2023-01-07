package gen

import (
	"context"

	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Accounts {
	arg := CreateAuthorParams{
		Owner:    util.RandonOwner(),
		Balance:  int32(util.RandomMoney()),
		Currency: util.RandomCurrency(),
	}
	out, err := testQueries.CreateAuthor(context.Background(), arg)
	if err != nil {
		require.NoError(t, err)
		require.NotEmpty(t, out)
		require.Equal(t, arg.Owner, out.Owner)
		require.Equal(t, arg.Balance, out.Balance)
		require.Equal(t, arg.Currency, out.Currency)
		require.NotZero(t, out.ID)
		require.NotZero(t, out.CreatedAt)
	}
	return out

}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)

}
func TestUpdateAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: account1.Balance,
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	if err != nil {
		require.Error(t, err)
		require.Equal(t, account2, account1)
	}

}
func TestDeleteAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	if err != nil {
		require.Error(t, err)
	}
}
