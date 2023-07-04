-- name: CreateTransfers :execresult
INSERT INTO TRANSFERS (from_account_id, to_account_id, amount) VALUE (?, ?, ?);

-- name: SelectLastIntroIdTransfer :one
select *
FROM TRANSFERS
WHERE id = last_insert_id();

-- name: ListTransfersFrom :many
SELECT *
FROM TRANSFERS
WHERE from_account_id = ?
LIMIT ? OFFSET ?;

-- name: ListTransfersTo :many
SELECT *
FROM TRANSFERS
WHERE to_account_id = ?
LIMIT ? OFFSET ?;

-- name: GetTransfer :one
SELECT *
FROM TRANSFERS
WHERE id = ?;