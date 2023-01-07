-- name: CreateAuthor :one
INSERT INTO accounts (
  owner,balance,
  currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreateTransfers :one
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount 
) VALUES (
  $1, $2,$3 
)
RETURNING *;

-- name: CreateEntries :one
INSERT INTO entries (
  account_id,
  amount 
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: GetTransfers :one
SELECT * FROM transfers WHERE id = $1 LIMIT 1;

-- name: ListAccount :many
SELECT * FROM accounts ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts SET balance = $2 WHERE id = $1 RETURNING *;

-- name: deleteAccount :exec
DELETE from accounts WHERE id = $1;

-- name: createTeacher :one 
INSERT INTO teacher (
  subject,
  salary
) VALUES (
  $1,$2
)
RETURNING *;






