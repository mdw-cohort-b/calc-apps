package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mdw-cohort-b/calc-apps/handlers"
)

func main() {
	router := handlers.NewHTTPRouter(os.Stderr)
	address := "localhost:8080"
	log.Println("Listening on " + address)
	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(err)
	}
}
