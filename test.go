/*
	This module is just for testing purposes

probing the behaviour of different codes
*/
package main

import (
	"encoding/csv"
	"os"
)

func main() {
	// Open the CSV file in append mode
	file, err := os.OpenFile("files/data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// Handle the error
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)

	// Write the data to the file
	writer.Write([]string{"John Doe", "or@example.com"})
	writer.Flush()
}
