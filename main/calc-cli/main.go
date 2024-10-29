package main

import (
	"log"
	"os"

	"github.com/mdw-cohort-b/calc-apps/handlers"
	"github.com/mdw-cohort-b/calc-lib"
)

func main() {
	handler := handlers.NewCLIHandler(os.Stdout, &calc.Addition{})
	err := handler.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
