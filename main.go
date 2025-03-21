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

func WriteToFile(filename string, expense string) {
	expenseFile, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("File cannot be opened: %v", err)
	}
	defer expenseFile.Close()

	_, err = expenseFile.WriteString(expense)
	if err != nil {
		log.Fatalf("Data cannot be written to file: %v", err)
	}
	fmt.Println("New expense added to file...")
}

func main() {
	file := "expenses.txt"
	CreateExpenseFile(file)
	WriteToFile(file, "Car: £331.11\n")
}
