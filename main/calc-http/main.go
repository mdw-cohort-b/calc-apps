package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	//router := http.NewServeMux()
	//router.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %s", r.URL.Query().Get("name"))
	//})

	logger := log.New(os.Stdout, "http: ", 0)
	router := handlers.NewHTTPRouter(logger)
	address := "localhost:8080"
	log.Println("Listening on " + address)
	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(err)
	}
}
