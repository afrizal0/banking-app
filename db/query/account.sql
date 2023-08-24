-- name: CreateAccount :one
INSERT INTO account (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;


-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM account
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccount :many
SELECT * FROM account
ORDER BY id 
LIMIT $1
OFFSET $2;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;

-- name: UpdateAccount :one
UPDATE account SET balance = $1
WHERE id = $2
RETURNING *;

-- name: AddAccountBalance :one
UPDATE account 
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

