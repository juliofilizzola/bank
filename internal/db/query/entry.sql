-- name: CreateEntry :execresult
INSERT INTO entries (account_id, amount) VALUE (?, ?);

-- name: SelectLastIntroIdEntry :one
select * FROM entries WHERE id = last_insert_id();

-- name: GetEntry :one
SELECT * FROM entries WHERE id = ?;

-- name: ListEntries :many
SELECT * FROM ENTRIES WHERE account_id = ? LIMIT ? OFFSET ?;