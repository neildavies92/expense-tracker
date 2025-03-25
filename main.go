package main

import (
	"bufio"
	"expense-tracker/csv"
	"fmt"
	"os"
)

func getUserInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	filename := "test.csv"

	writer, file, err := csv.CreateCSV(filename)
	if err != nil {
		fmt.Printf("Error creating CSV file: %v\n", err)
		return
	}
	defer file.Close()

	for {
		expenseName := getUserInput("Enter name of expense (or type 'done' to finish): ")
		if expenseName == "done" {
			break
		}
		expenseValue := (getUserInput("Enter value of expense: "))
		dueDate := getUserInput("Enter due date of expense: ")

		record := []string{expenseName, expenseValue, dueDate}
		csv.AppendExpenses(writer, record)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing CSV writer: ", err)
	}
}
