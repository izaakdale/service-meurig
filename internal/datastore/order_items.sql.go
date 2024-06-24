// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: order_items.sql

package datastore

import (
	"context"
)

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO order_items (
  id,
  order_id, 
  product_id,
  product_name,
  quantity
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, order_id, product_id, product_name, quantity
`

type CreateOrderItemParams struct {
	ID          string `json:"id"`
	OrderID     string `json:"order_id"`
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int32  `json:"quantity"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRow(ctx, createOrderItem,
		arg.ID,
		arg.OrderID,
		arg.ProductID,
		arg.ProductName,
		arg.Quantity,
	)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.ProductID,
		&i.ProductName,
		&i.Quantity,
	)
	return i, err
}

const getOrderItems = `-- name: GetOrderItems :many
SELECT id, order_id, product_id, product_name, quantity FROM order_items
WHERE order_id = $1
`

func (q *Queries) GetOrderItems(ctx context.Context, orderID string) ([]OrderItem, error) {
	rows, err := q.db.Query(ctx, getOrderItems, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrderItem
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.ProductID,
			&i.ProductName,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}