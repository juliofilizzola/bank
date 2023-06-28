-- name: CreateTransfers :exec
INSERT INTO transfers (from_account_id, to_account_id, amount) VALUE (?, ?, ?);

-- name: ListTransfersFrom :many
SELECT * FROM transfers WHERE from_account_id = ? LIMIT ? OFFSET ?;

-- name: ListTransfersTo :many
SELECT * FROM transfers WHERE to_account_id = ? LIMIT ? OFFSET ?;
