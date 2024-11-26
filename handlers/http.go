package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mdw-cohort-b/calc-lib"
)

func NewHTTPRouter(logger *log.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /add", NewHTTPHandler(logger, new(calc.Addition)))
	mux.Handle("GET /sub", NewHTTPHandler(logger, new(calc.Subtraction)))
	mux.Handle("GET /mul", NewHTTPHandler(logger, new(calc.Multiplication)))
	mux.Handle("GET /div", NewHTTPHandler(logger, new(calc.Division)))
	return mux
}

type HTTPHandler struct {
	logger     *log.Logger
	calculator Calculator
}

func NewHTTPHandler(logger *log.Logger, calculator Calculator) *HTTPHandler {
	return &HTTPHandler{logger: logger, calculator: calculator}
}
func (this *HTTPHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	rawA := query.Get("a")
	rawB := query.Get("b")
	a, err := strconv.Atoi(rawA)
	if err != nil {
		http.Error(response, "The a parameter must be an integer", http.StatusUnprocessableEntity)
		return
	}
	b, err := strconv.Atoi(rawB)
	if err != nil {
		http.Error(response, "The b parameter must be an integer", http.StatusUnprocessableEntity)
		return
	}
	c := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintf(response, "%d", c)
	if err != nil {
		this.logger.Println(err)
	}
}
