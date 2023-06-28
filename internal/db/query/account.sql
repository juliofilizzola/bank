-- name: CreateAccount :one
INSERT INTO accounts (owner, balance, currency) VALUES (?, ?, ?) SELECT;

-- name: GetAccount :one
SELECT * FROM accounts WHERE id = ? LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id LIMIT ? OFFSET ?;

-- name: UpdatedAccounts :one
UPDATE accounts SET balance = ? where id = ?;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = ?;