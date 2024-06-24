-- name: CreateProduct :one
INSERT INTO products (
  id,
  title, 
  description, 
  price
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY id
LIMIT $1
OFFSET $2;