package handlers

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strconv"
)

type CSVHandler struct {
	logger      *log.Logger
	input       *csv.Reader
	output      *csv.Writer
	calculators map[string]Calculator
}

func NewCSVHandler(logger *log.Logger, input io.Reader, output io.Writer,
	calculators map[string]Calculator) *CSVHandler {

	return &CSVHandler{
		logger:      logger,
		input:       csv.NewReader(input),
		output:      csv.NewWriter(output),
		calculators: calculators,
	}
}
func (this *CSVHandler) Handle() error {
	this.input.FieldsPerRecord = 3
	defer this.output.Flush()
	for {
		record, err := this.input.Read()
		if err == io.EOF {
			break
		}
		if errors.Is(err, csv.ErrFieldCount) {
			this.logger.Println(err)
			continue
		}
		// TODO: if err != nil
		a, err := strconv.Atoi(record[0])
		if err != nil {
			this.logger.Println("invalid arg:", record[0])
			continue
		}
		b, err := strconv.Atoi(record[2])
		if err != nil {
			this.logger.Println("invalid arg:", record[2])
			continue
		}
		calculator, ok := this.calculators[record[1]]
		if !ok {
			this.logger.Println("unsupported operator:", record[1])
			continue
		}
		c := calculator.Calculate(a, b)
		_ = this.output.Write(append(record, strconv.Itoa(c)))
		// TODO: if err != nil...
	}
	return this.output.Error()
}