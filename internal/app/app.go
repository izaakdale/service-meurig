package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/izaakdale/service-meurig/internal/datastore"
	"github.com/izaakdale/service-meurig/internal/router"
	"github.com/jackc/pgx/v5"
)

func Run() error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_CONNECTION_STRING"))
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	q := datastore.New(conn)

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
		Handler: router.New(q),
	}
	errCh := make(chan error)
	go func(ch chan error) {
		ch <- srv.ListenAndServe()
	}(errCh)

	shutCh := make(chan os.Signal, 1)
	signal.Notify(shutCh, os.Interrupt, syscall.SIGTERM)

	log.Println("server started")
	select {
	case <-shutCh:
		return srv.Shutdown(ctx)
	case err := <-errCh:
		return fmt.Errorf("server errored: %w", err)
	}
}
