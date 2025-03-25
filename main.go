package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func createCSV(filename string) (*csv.Writer, *os.File, error) {
	expenseFile, err := os.Create(filename)
	if err != nil {
		log.Fatalf("File cannot be created: %v", err)
	}
	writer := csv.NewWriter(expenseFile)
	return writer, expenseFile, nil
}

func writeToCSV(writer *csv.Writer, record []string) {
	err := writer.Write(record)
	if err != nil {
		fmt.Println("Error writing record to CSV: %v", err)
	}
}

func main() {
	filename := "test.csv"

	writer, file, err := createCSV(filename)
	if err != nil {
		fmt.Println("Error creating CSV file: %v", err)
		return
	}
	defer file.Close()

	header := []string{"Expense", "Amount", "Due Date"}
	writeToCSV(writer, header)

	records := [][]string{
		{"Rent", "1000", "1st"},
		{"Car", "350", "1st"},
	}

	for _, record := range records {
		writeToCSV(writer, record)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing CSV writer: ", err)
	}
}
