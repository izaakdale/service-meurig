package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/izaakdale/service-meurig/internal/datastore"
)

type itemInserter interface {
	CreateOrderItem(ctx context.Context, params datastore.CreateOrderItemParams) (datastore.OrderItem, error)
}

type insertItemRequest struct {
	ProductID   string `json:"product_id,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	Quantity    int32  `json:"quantity,omitempty"`
}

func addItem(inserter itemInserter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req insertItemRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		params := datastore.CreateOrderItemParams{
			ID:          uuid.NewString(),
			OrderID:     r.PathValue(orderIDPathValue),
			ProductID:   req.ProductID,
			ProductName: req.ProductName,
			Quantity:    req.Quantity,
		}

		item, err := inserter.CreateOrderItem(r.Context(), params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err = json.NewEncoder(w).Encode(item); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type itemsFetcher interface {
	GetOrderItems(ctx context.Context, orderID string) ([]datastore.OrderItem, error)
}

func getOrderItems(fetcher itemsFetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := fetcher.GetOrderItems(r.Context(), r.PathValue(orderIDPathValue))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(items); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
