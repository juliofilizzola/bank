-- name: CreateEntry :execresult
INSERT INTO ENTRIES (account_id, amount) VALUE (?, ?);

-- name: SelectLastIntroIdEntry :one
select *
FROM ENTRIES
WHERE id = last_insert_id();

-- name: GetEntry :one
SELECT *
FROM ENTRIES
WHERE id = ?;

-- name: ListEntries :many
SELECT *
FROM ENTRIES
WHERE account_id = ?
LIMIT ? OFFSET ?;