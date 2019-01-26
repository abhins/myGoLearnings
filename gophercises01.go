package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"io"
)

func main() {
    file, err := os.Open("problems.csv")
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file"))
	}
	r := csv.NewReader(file)

   for {
		record, err := r.Read()
//		fmt.Println(err)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
