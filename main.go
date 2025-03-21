package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("expenses.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("File cannot be created: %v", err)
	}
	defer file.Close()

	if _, err := os.Stat("expenses.txt"); os.IsNotExist(err) {
		fmt.Println("File succesfully created.")
	} else {
		fmt.Println("Skipped creating... file already exists.")
	}
}
