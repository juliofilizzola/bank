-- name: CreateAccount :exec
INSERT INTO accounts (balance, owner, currency) VALUES (?, ?, ?);

-- name: GetAccount :one
SELECT * FROM accounts WHERE id = ? LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id LIMIT ? OFFSET ?;

-- name: UpdatedAccounts :exec
UPDATE accounts SET balance = ? where id = ?;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = ?;