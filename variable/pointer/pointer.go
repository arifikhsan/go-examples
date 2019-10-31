package main

import (
	"fmt"
)

func zeroval(initialValue int) {
	initialValue = 0 // has no effect or copied
}

func zeroPointer(initialPointer *int) {
	*initialPointer = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroPointer(&i)
	fmt.Println("zero pointer: ", i)

	fmt.Println("pointer: ", &i)
}