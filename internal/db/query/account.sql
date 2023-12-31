-- name: CreateAccount :execresult
INSERT INTO accounts (balance, owner, currency) VALUES (?, ?, ?);

-- name: SelectLastIntroID :one
select * FROM accounts WHERE id = last_insert_id();

-- name: GetAccount :one
SELECT * FROM accounts WHERE id = sqlc.arg(id) LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id LIMIT ? OFFSET ?;

-- name: AddBalanceUser :exec
UPDATE accounts SET balance = balance + sqlc.arg(amount) where id = sqlc.arg(id);

-- name: RemoveBalanceUser :exec
UPDATE accounts SET balance = balance - sqlc.arg(amount) where id = sqlc.arg(id);

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = sqlc.arg(id);