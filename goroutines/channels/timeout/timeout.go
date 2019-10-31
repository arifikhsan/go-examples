package main

import (
	"fmt"
	"time"
)

func main() {
	channelOne := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		channelOne <- "result 1"
	}()

	select {
	case result := <-channelOne:
		fmt.Println(result)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	channelTwo := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		channelTwo <- "result 2"
	}()
	select {
	case result := <-channelTwo:
		fmt.Println(result)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
