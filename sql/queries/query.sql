-- name: CreateUser :exec
INSERT INTO users (username,email) VALUES ($1,$2) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1;

-- name: CreateAccount :exec
INSERT INTO accounts (id,owner,balance) VALUES ($1,$2,$3);

-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: UpdateAccount :exec
UPDATE accounts SET balance = $2 WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

-- name: CreateEntry :exec
INSERT INTO entries (id,account_id,transaction_type,amount) VALUES ($1,$2,$3,$4);

-- name: BulkCreateEntry :exec
INSERT INTO entries (id,account_id,transaction_type,amount) values($1,$2,$3,$4),($5,$6,$7,$8);

-- name: GetEntry :one
SELECT * FROM entries WHERE id = $1 LIMIT 1;

-- name: ListEntry :many
SELECT * FROM entries WHERE account_id = $1 ORDER BY id LIMIT $2 OFFSET $3;

-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = $1;

-- name: CreateTransfer :exec
INSERT INTO transfers (id,from_account_id,to_account_id,amount) VALUES ($1,$2,$3,$4);

-- name: GetTransfer :one
SELECT * FROM transfers WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = $1;