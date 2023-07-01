// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: entry.sql

package db

import (
	"context"
	"database/sql"
)

const createEntry = `-- name: CreateEntry :execresult
INSERT INTO entries (account_id, amount) VALUE (?, ?)
`

type CreateEntryParams struct {
	AccountID int32
	Amount    int64
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createEntry, arg.AccountID, arg.Amount)
}

const getEntry = `-- name: GetEntry :one
SELECT id, account_id, amount, created_at FROM entries WHERE id = ?
`

func (q *Queries) GetEntry(ctx context.Context, id int32) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, account_id, amount, created_at FROM ENTRIES WHERE account_id = ? LIMIT ? OFFSET ?
`

type ListEntriesParams struct {
	AccountID int32
	Limit     int32
	Offset    int32
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectLastIntroIdEntry = `-- name: SelectLastIntroIdEntry :one
select id, account_id, amount, created_at FROM entries WHERE id = last_insert_id()
`

func (q *Queries) SelectLastIntroIdEntry(ctx context.Context) (Entry, error) {
	row := q.db.QueryRowContext(ctx, selectLastIntroIdEntry)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
