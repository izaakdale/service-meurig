package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/izaakdale/ittp"
	"github.com/izaakdale/service-meurig/internal/datastore"
)

var (
	// customerIDPathValue = "customerID"
	orderIDPathValue = "orderID"
)

type dbClient interface {
	CreateCustomer(ctx context.Context, params datastore.CreateCustomerParams) (datastore.Customer, error)

	ListProducts(ctx context.Context, params datastore.ListProductsParams) ([]datastore.Product, error)
	CreateProduct(ctx context.Context, params datastore.CreateProductParams) (datastore.Product, error)

	GetCustomerOrders(ctx context.Context, customerID string) ([]datastore.Order, error)
	CreateOrder(ctx context.Context, params datastore.CreateOrderParams) (datastore.Order, error)
	UpdateOrderStatus(ctx context.Context, params datastore.UpdateOrderStatusParams) (datastore.Order, error)

	CreateOrderItem(ctx context.Context, params datastore.CreateOrderItemParams) (datastore.OrderItem, error)
	GetOrderItems(ctx context.Context, orderID string) ([]datastore.OrderItem, error)
}

func New(cli dbClient) http.Handler {
	mux := ittp.NewServeMux()

	mux.Post("/customers", createCustomer(cli))

	mux.Get("/products", getProductsList(cli))
	mux.Post("/products", insertProduct(cli))

	mux.Get("/orders", getCustomerOrder(cli))
	mux.Post("/orders", createOrder(cli))

	mux.Put(fmt.Sprintf("/orders/{%s}", orderIDPathValue), updateOrderStatus(cli))
	mux.Post(fmt.Sprintf("/orders/{%s}", orderIDPathValue), addItem(cli))
	mux.Get(fmt.Sprintf("/orders/{%s}", orderIDPathValue), getOrderItems(cli))

	return mux
}
