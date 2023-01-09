package store

import (
	"context"
	"database/sql"
	"simplebank/db/gen"
	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	drivername   = "postgresq12"
	driverSource = "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable"
)

func TestTransferTx(t *testing.T) {
	dbcon, err := sql.Open(drivername, driverSource)
	if err == nil {
		s := newStore(dbcon)
		account1 := gen.CreateAuthorParams{

			Owner:    util.RandomOwner(),
			Balance:  int32(util.RandomMoney()),
			Currency: util.RandomCurrency(),
		}
		account2 := gen.CreateAuthorParams{

			Owner:    util.RandomOwner(),
			Balance:  int32(util.RandomMoney()),
			Currency: util.RandomCurrency(),
		}
		n := 5
		amount := 10
		errChan := make(chan error)
		tranParm := make(chan TransferTxResult)

		for i := 0; i < n; i++ {
			go func() {
				res, err := s.TransferTx(context.Background(), TransferTxParams{
					FromAccountID: 1,
					ToAccountID:   2,
					Amount:        int32(amount),
				})
				errChan <- err
				tranParm <- res
			}()
		}
		for i := 0; i < n; i++ {
			err := <-errChan
			require.NoError(t, err)
			result := <-tranParm
			transfer := result.Transfer
			require.NotEmpty(t, transfer)
			require.Equal(t, 1, transfer.FromAccountID)
			require.Equal(t, 2, transfer.ToAccountID)
			require.Equal(t, amount, transfer.Amount)
			require.NotZero(t, transfer.ID)
			require.NotZero(t, transfer.CreatedAt)

			_, err1 := s.GetTransfers(context.Background(), transfer.ID)
			require.NoError(t, err1)
			// check entries
			fromEntry := result.FromEntry
			require.NotEmpty(t, fromEntry)
			require.Equal(t, 1, fromEntry.AccountID)
			require.Equal(t, -amount, fromEntry.Amount)
			require.NotZero(t, fromEntry.ID)
			require.NotZero(t, fromEntry.CreatedAt)
			_, err2 := s.GetEntry(context.Background(), fromEntry.ID)
			require.NoError(t, err2)
			toEntry := result.ToEntry
			require.NotEmpty(t, toEntry)
			require.Equal(t, 2, toEntry.AccountID)
			require.Equal(t, amount, toEntry.Amount)
			require.NotZero(t, toEntry.ID)
			require.NotZero(t, toEntry.CreatedAt)
			_, err = s.GetEntry(context.Background(), toEntry.ID)
			require.NoError(t, err)
			// TODO: check accounts' balance
			fromAccount := result.FromAccount
			require.NotEmpty(t, fromAccount)
			require.Equal(t, 1, fromAccount.ID)
			toAccount := result.ToAccount
			require.NotEmpty(t, toAccount)
			require.Equal(t, 2, toAccount.ID)
			// check accounts' balance
			diff1 := account1.Balance - fromAccount.Balance
			diff2 := toAccount.Balance - account2.Balance
			require.Equal(t, diff1, diff2)
			require.True(t, diff1 > 0)
			require.True(t, diff1%int32(amount) == 0) // 1 * amount, 2 * amount, 3* amount, ...., n * amount
			k := int(diff1 / int32(amount))
			require.True(t, k >= 1 && k <= n)

		}
	}

}
