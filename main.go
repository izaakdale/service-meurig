package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	commerce "github.com/izaakdale/service-meurig/api/commerce/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	commerce.StoreServer
}

// FetchProducts implements commerce.StoreServer.
func (s *server) FetchProducts(ctx context.Context, req *commerce.ProductListRequest) (*commerce.ProductListResponse, error) {
	return &commerce.ProductListResponse{
		Products: []*commerce.Product{
			{
				Id:          "0001",
				Name:        "item-1",
				Description: "some item for sale",
				Price:       99.99,
				ImageUrl:    "fake-url.abc",
				Aisle:       "4",
			},
			{
				Id:          "0002",
				Name:        "item-2",
				Description: "some item for sale",
				Price:       99.99,
				ImageUrl:    "fake-url.abc",
				Aisle:       "2",
			},
			{
				Id:          "0003",
				Name:        "item-3",
				Description: "some item for sale",
				Price:       99.99,
				ImageUrl:    "fake-url.abc",
				Aisle:       "7",
			},
		},
	}, nil
}

func main() {
	ls, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
	if err != nil {
		panic(err)
	}

	gsrv := grpc.NewServer()
	reflection.Register(gsrv)

	srv := server{}

	commerce.RegisterStoreServer(gsrv, &srv)

	errCh := make(chan error)
	go func(ch chan error) {
		ch <- gsrv.Serve(ls)
	}(errCh)

	shCh := make(chan os.Signal, 2)
	signal.Notify(shCh, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-shCh:
			os.Exit(1)
		case err := <-errCh:
			log.Fatalf("grpc server errored: %v", err)
		}
	}
}
