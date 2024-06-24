// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package datastore

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Customer struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type Order struct {
	ID         string           `json:"id"`
	CustomerID string           `json:"customer_id"`
	Status     string           `json:"status"`
	Total      int32            `json:"total"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

type OrderItem struct {
	ID          string `json:"id"`
	OrderID     string `json:"order_id"`
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int32  `json:"quantity"`
}

type Product struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Price       int32            `json:"price"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}
