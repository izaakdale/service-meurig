package main

import (
	"log"

	"github.com/izaakdale/service-meurig/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
