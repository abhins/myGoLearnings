package main

import "fmt"

func ping(c chan string) {
i := 0
	for i < 1 {
		c <- "ping"
	}

}

func pong(c chan string) {
i := 0
	for i < 1 {
		c <- "pong"
	}

}

func main() {
    c := make(chan string,1)
	go ping(c)
	go pong(c)
	for {
		x,y := <-c, <-c
		fmt.Println(x, y)
	}
}
