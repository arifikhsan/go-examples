package main

import (
	"fmt"
	"time"
)

func main() {
	channelOne := make(chan string)
	channelTwo := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channelOne <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channelTwo <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case messageOne := <-channelOne:
			fmt.Println("received", messageOne)
		case messageTwo := <-channelTwo:
			fmt.Println("received", messageTwo)
		}
	}
}
