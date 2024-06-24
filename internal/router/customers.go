package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/izaakdale/service-meurig/internal/datastore"
)

type customerInserter interface {
	CreateCustomer(ctx context.Context, params datastore.CreateCustomerParams) (datastore.Customer, error)
}

func createCustomer(inserter customerInserter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customer, err := inserter.CreateCustomer(r.Context(), datastore.CreateCustomerParams{
			ID:       "stub",
			Username: "stub",
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err = json.NewEncoder(w).Encode(customer); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
