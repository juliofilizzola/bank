-- name: CreateTransfers :execresult
INSERT INTO transfers (from_account_id, to_account_id, amount) VALUE (?, ?, ?);

-- name: SelectLastIntroIdTransfer :one
select * FROM transfers WHERE id = last_insert_id();

-- name: ListTransfersFrom :many
SELECT * FROM transfers WHERE from_account_id = ? LIMIT ? OFFSET ?;

-- name: ListTransfersTo :many
SELECT * FROM transfers WHERE to_account_id = ? LIMIT ? OFFSET ?;

-- name: GetTransfer :one
SELECT * FROM TRANSFERS WHERE id = ?;