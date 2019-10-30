// channel are the pipes that connect concurrent goroutines
package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	mgs := <- messages
	fmt.Println(mgs)
}
