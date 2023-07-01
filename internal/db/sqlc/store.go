package db

import (
	"context"
	"database/sql"
	"errors"
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
	var ctxB = context.Background()
	account, err := s.GetAccount(ctxB, arg.FromAccountID)
	account2, err := s.GetAccount(ctxB, arg.ToAccountId)
	if err != nil {
		return result, err
	}

	if account.Balance < arg.Amount {
		fmt.Println("account not balance")
		return result, errors.New("account not balance from transaction")
	}

	err = s.execTx(ctx, func(queries *Queries) error {
		var err error

		res, err := queries.CreateTransfers(ctx, CreateTransfersParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}

		var idTransfer = int(id)
		var idConvert = int32(idTransfer)

		result.Transfer, err = queries.GetTransfer(ctxB, idConvert)

		if err != nil {
			return err
		}

		_, err = queries.CreateEntry(ctxB, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = queries.SelectLastIntroIdEntry(ctxB)

		if err != nil {
			return err
		}

		_, err = queries.CreateEntry(ctxB, CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = queries.SelectLastIntroIdEntry(ctxB)

		if err != nil {
			return err
		}

		err = addMoney(ctx, queries, account2.ID, arg.Amount)

		if err != nil {
			return err
		}

		err = removeMoney(ctx, queries, account.ID, arg.Amount)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return result, err
	}
	return result, nil
}

func addMoney(ctx context.Context, q *Queries, id int32, amount int64) error {
	err := q.AddBalanceUser(ctx, AddBalanceUserParams{
		Amount: amount,
		ID:     id,
	})

	if err != nil {
		return err
	}
	return nil
}

func removeMoney(ctx context.Context, q *Queries, id int32, amount int64) error {
	err := q.RemoveBalanceUser(ctx, RemoveBalanceUserParams{
		Amount: amount,
		ID:     id,
	})

	if err != nil {
		return err
	}

	return nil
}
