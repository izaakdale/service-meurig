
-- name: CreateOrder :one
INSERT INTO orders (
  id,
  customer_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateOrderTotalPrice :one
UPDATE orders
SET total = total + $2
WHERE id = $1
RETURNING *;

-- name: UpdateOrderStatus :one
UPDATE orders
SET status = $2
WHERE id = $1
RETURNING *;

-- name: GetCustomerOrders :many
SELECT * FROM orders
WHERE customer_id = $1;