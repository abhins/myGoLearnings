package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}

}

func endless() {
i := 0
	for i < 1 {
		fmt.Println("in infinite\n")
		time.Sleep(100 * time.Millisecond)
	}

}

func main() {
	go endless()
	go say("world")
	say("hello")
}
