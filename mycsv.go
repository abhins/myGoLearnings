package main

import (
	"encoding/csv"
	"os"
	"io"
	"fmt"
	"strconv"
)

func main() {
//read from a csv file
	var r io.Reader
	total  := 0
	correct := 0
	r, err := os.Open("problems.csv")
	if err != nil {
		os.Exit(3)
	}
	t := csv.NewReader(r)
    for {
		rec, err := t.Read()
		    if err == io.EOF {
			    fmt.Printf("Your Quiz Score: %d/%d\n", correct,total)
				break
			}
			if err != nil {
				os.Exit(4)
			}
			total = total + 1
            resp := getResponse(rec)
			if resp == true {
				correct = correct + 1
			}
	}
}

func getResponse(b []string) bool {
		var ans int
		fmt.Printf("What's %s?\n", b[0])
		fmt.Scan(&ans)
		i, err := strconv.Atoi(b[1])
		if err != nil {
			os.Exit(4)
		}
		if ans == i {
			return true
		} else {
			return false
		}
}
