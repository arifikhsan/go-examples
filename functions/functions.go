package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	response := plus(1, 2)
	fmt.Println("1 + 2 =", response)

	response = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", response)
}
