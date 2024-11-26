package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mdw-cohort-b/calc-apps/handlers"
)

func main() {
	logger := log.New(os.Stdout, "http: ", 0)
	router := handlers.NewHTTPRouter(logger)
	address := "localhost:8080"
	log.Println("Listening on " + address)
	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(err)
	}
}
