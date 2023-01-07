package store

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/db/gen"
	_ "simplebank/util"
)

type store struct {
	*gen.Queries
	db *sql.DB
}

func newStore(db *sql.DB) *store {
	return &store{
		db:      db,
		Queries: gen.New(db),
	}
}

// execTX executes a function within a database connection
func (s *store) execTX(ctx context.Context, fn func(*gen.Queries) error) error {
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
	Transfer    gen.Transfers `json: "transfer"`
	FromAccount gen.Accounts  `json: "from_account"`
	ToAccount   gen.Accounts  `json: "to_account"`
	FromEntry   gen.Entries   `json: "from_entry"`
	ToEntry     gen.Entries   `json: "to_entry"`
}

func (s *store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
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

		return nil
	})
	return res, err

}
