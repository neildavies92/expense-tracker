package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func CreateCSV(filename string) (*csv.Writer, *os.File, error) {
	_, err := os.Stat(filename)
	isNewFile := os.IsNotExist(err)

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("File cannot be created: %v", err)
	}
	writer := csv.NewWriter(file)

	if isNewFile {
		header := []string{"Expense", "Amount", "Due Date"}
		if err := writer.Write(header); err != nil {
			log.Fatalf("Error writing header to CSV: %v", err)
		}
	}
	return writer, file, nil
}

func AppendExpenses(writer *csv.Writer, record []string) {
	err := writer.Write(record)
	if err != nil {
		fmt.Printf("Error writing record to CSV: %v\n", err)
	}
}
