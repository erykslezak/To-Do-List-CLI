package main

import (
	"encoding/csv"
	"fmt"
	"github.com/mergestat/timediff"
	"log"
	"os"
	"text/tabwriter"
	"time"
)

func main() {
	PrintToDo()
}

func PrintToDo() {
	// Open the CSV file
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader
	r := csv.NewReader(f)
	layout := time.RFC3339
	// Read all records from the CSV file
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	// Iterate through the records and print each one
	for _, record := range records {
		recordTime, _ := time.Parse(layout, record[2])
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", record[0], record[1], timediff.TimeDiff(recordTime), record[3])
	}

	// Flush the writer to ensure all data is written to the output
	w.Flush()
}
