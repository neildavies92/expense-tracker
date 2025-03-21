package main

import (
	"fmt"
	"log"
	"os"
)

func CreateExpenseFile(filename string) {
	expenseFile, err := os.OpenFile(filename, os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("File cannot be created: %v", err)
	}
	defer expenseFile.Close()
	fmt.Println("File created successfully...")
}

func WriteToFile(filename string, data string) {
	expenseFile, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("File cannot be opened: %v", err)
	}
	defer expenseFile.Close()

	_, err = expenseFile.WriteString(data)
	if err != nil {
		log.Fatalf("Data cannot be written to file: %v", err)
	}
	fmt.Println("Data written successfully...")
}

func main() {
	CreateExpenseFile("expenses.txt")
	WriteToFile("expenses.txt", "Hello World\n")
}
