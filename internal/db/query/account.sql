-- name: CreateAccount :exec
insert into accounts (owner, balance, currency) values (?, ?, ?);