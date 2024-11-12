package main

import (
	"log"
	"os"

	"github.com/mdw-cohort-b/calc-apps/handlers"
	"github.com/mdw-cohort-b/calc-lib"
)

func main() {
	logger := log.New(os.Stderr, ">>> ", 0)
	handler := handlers.NewCSVHandler(logger, os.Stdin, os.Stdout, calculators)
	err := handler.Handle()
	if err != nil {
		logger.Fatal(err)
	}
}

var calculators = map[string]handlers.Calculator{
	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
}
