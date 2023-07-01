package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

type TransferTxParams struct {
	FromAccountID int32 `json:"from_account_id"`
	ToAccountId   int32 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (s Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := s.execTx(ctx, func(queries *Queries) error {
		var err error
		_, err = queries.CreateTransfers(ctx, CreateTransfersParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		result.Transfer, err = queries.SelectLastIntroIdTransfer(context.Background())

		if err != nil {
			return err
		}

		_, err = queries.CreateEntry(context.Background(), CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = queries.SelectLastIntroIdEntry(context.Background())
		if err != nil {
			return err
		}

		_, err = queries.CreateEntry(context.Background(), CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = queries.SelectLastIntroIdEntry(context.Background())
		if err != nil {
			return err
		}
		// todo: update balance

		return nil
	})

	if err != nil {
		return result, err
	}
	return result, nil
}
