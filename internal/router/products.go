package router

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/izaakdale/service-meurig/internal/datastore"
)

type productLister interface {
	ListProducts(ctx context.Context, params datastore.ListProductsParams) ([]datastore.Product, error)
}

type listProductsRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func getProductsList(lister productLister) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req listProductsRequest

		limitStr := r.URL.Query().Get("limit")
		offsetStr := r.URL.Query().Get("offset")

		if limitStr != "" {
			limit, err := strconv.Atoi(limitStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			req.Limit = int32(limit)
		} else {
			req.Limit = 10
		}

		if offsetStr != "" {
			offset, err := strconv.Atoi(offsetStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			req.Offset = int32(offset)
		} else {
			req.Offset = 0
		}

		products, err := lister.ListProducts(r.Context(), datastore.ListProductsParams{
			Limit:  req.Limit,
			Offset: req.Offset,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type productInserter interface {
	CreateProduct(ctx context.Context, params datastore.CreateProductParams) (datastore.Product, error)
}

func insertProduct(inserter productInserter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var params datastore.CreateProductParams

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		params.ID = uuid.NewString()
		product, err := inserter.CreateProduct(r.Context(), params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err = json.NewEncoder(w).Encode(product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
