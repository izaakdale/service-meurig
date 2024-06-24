package router

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/izaakdale/service-meurig/internal/datastore"
)

type orderFetcher interface {
	GetCustomerOrders(ctx context.Context, customerID string) ([]datastore.Order, error)
}

func getCustomerOrder(fetcher orderFetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get customerID from somewhere
		customerID := "stub"
		order, err := fetcher.GetCustomerOrders(r.Context(), customerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(order); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type orderCreator interface {
	CreateOrder(ctx context.Context, params datastore.CreateOrderParams) (datastore.Order, error)
}

func createOrder(creator orderCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		order, err := creator.CreateOrder(r.Context(), datastore.CreateOrderParams{
			ID:         uuid.NewString(),
			CustomerID: "stub",
			// get customerID from somewhere
			// CustomerID: r.PathValue("customerID"),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err = json.NewEncoder(w).Encode(order); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type statusUpdater interface {
	UpdateOrderStatus(ctx context.Context, params datastore.UpdateOrderStatusParams) (datastore.Order, error)
}

type updateStatusRequest struct {
	Status string `json:"status"`
}

func updateOrderStatus(updater statusUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderID := r.PathValue(orderIDPathValue)

		var req updateStatusRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		req.Status = strings.ToLower(req.Status)

		if req.Status == "" || !validateStatus(req.Status) {
			http.Error(w, "invalid status", http.StatusBadRequest)
			return
		}

		_, err := updater.UpdateOrderStatus(r.Context(), datastore.UpdateOrderStatusParams{
			ID:     orderID,
			Status: req.Status,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func validateStatus(status string) bool {
	return status == "cart" || status == "paid" || status == "shipped" || status == "cancelled"
}
