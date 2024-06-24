-- name: CreateOrderItem :one
INSERT INTO order_items (
  id,
  order_id, 
  product_id,
  product_name,
  quantity
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetOrderItems :many
SELECT * FROM order_items
WHERE order_id = $1;
