-- name: CreateCustomer :one
INSERT INTO customers (
  id,
  username
) VALUES (
  $1, $2
) RETURNING *;