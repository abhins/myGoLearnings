package main

import (
	"io"
	"os"
	"fmt"
)

func main() {
	var ii io.Reader 
	f,err := os.Open("newfile")
	if err != nil {
		//some error
	}
	ii = f
	b := make([]byte, 10)
	n, err := ii.Read(b)
	if err != nil {
		//some error
	}
	fmt.Printf("%d %s",n,b)
}
