package store

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/db/gen"
	_ "simplebank/util"
)

type Store interface {
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	gen.Querier
}

type SqlStore struct {
	*gen.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SqlStore{
		db:      db,
		Queries: gen.New(db),
	}
}

// execTX executes a function within a database connection
func (s *SqlStore) execTX(ctx context.Context, fn func(*gen.Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := gen.New(tx)
	err = fn(q)
	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			return fmt.Errorf("tx error: %v, roolbackerror is :%v", err, rberr)
		}
		return err
	}
	return tx.Commit()

}

type TransferTxParams struct {
	FromAccountID int32 `json:"from_account_id"`
	ToAccountID   int32 `json:"to_account_id"`
	Amount        int32 `json: "amount"`
}

type TransferTxResult struct {
	Transfer    gen.Transfer `json: "transfer"`
	FromAccount gen.Account  `json: "from_account"`
	ToAccount   gen.Account  `json: "to_account"`
	FromEntry   gen.Entry    `json: "from_entry"`
	ToEntry     gen.Entry    `json: "to_entry"`
}

func (s *SqlStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var res TransferTxResult
	var err1 error

	err := s.execTX(ctx, func(q *gen.Queries) error {

		res.Transfer, err1 = q.CreateTransfers(ctx, gen.CreateTransfersParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err1 != nil {
			return err1
		}
		res.FromEntry, err1 = q.CreateEntries(ctx, gen.CreateEntriesParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err1 != nil {
			return err1
		}
		res.ToEntry, err1 = q.CreateEntries(ctx, gen.CreateEntriesParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		// to do updating an account
		account1, err2 := s.GetAccount(ctx, int64(arg.FromAccountID))
		if err2 != nil {
			return err2

		}
		res.FromAccount, err1 = q.UpdateAccount(ctx, gen.UpdateAccountParams{
			ID:      int64(arg.FromAccountID),
			Balance: account1.Balance - arg.Amount,
		})
		if err1 != nil {
			return err1
		}
		account2, err3 := s.GetAccount(ctx, int64(arg.ToAccountID))
		if err3 != nil {
			return err3

		}

		res.ToAccount, err1 = s.UpdateAccount(ctx, gen.UpdateAccountParams{
			ID:      int64(arg.ToAccountID),
			Balance: account2.Balance + arg.Amount,
		})

		return nil
	})
	return res, err

}
