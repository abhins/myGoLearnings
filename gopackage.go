package main

import (
	"os"
	"log"
	"fmt"
)


func main() {

	file, err := os.Open("file.go") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 800)
	n, err := file.Read(b)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("%d %s",n,b)
	s := os.Environ()
	fmt.Printf("%q",s)
}
