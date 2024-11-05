package main

import (
	"flag"
	"log"
	"os"

	"github.com/mdw-cohort-b/calc-apps/handlers"
	"github.com/mdw-cohort-b/calc-lib"
)

func main() {
	var operation string
	flag.StringVar(&operation, "op", "+", "Operation to calculate")
	flag.Parse()

	handler := handlers.NewCLIHandler(os.Stdout, calculators[operation])
	err := handler.Handle(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
}

var calculators = map[string]handlers.Calculator{
	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
}
