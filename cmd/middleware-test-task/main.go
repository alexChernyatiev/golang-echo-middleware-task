package main

import (
	"log"

	"github.com/alexChernyatiev/golang-echo-middleware-task/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}

}
