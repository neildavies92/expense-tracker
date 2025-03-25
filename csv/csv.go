package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func CreateCSV(filename string) (*csv.Writer, *os.File, error) {
	expenseFile, err := os.Create(filename)
	if err != nil {
		log.Fatalf("File cannot be created: %v", err)
	}
	writer := csv.NewWriter(expenseFile)
	return writer, expenseFile, nil
}

func CreateCSVWriter(writer *csv.Writer, record []string) {
	err := writer.Write(record)
	if err != nil {
		fmt.Printf("Error writing record to CSV: %v\n", err)
	}
}
