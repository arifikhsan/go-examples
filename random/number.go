package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	sourceOne := rand.NewSource(time.Now().UnixNano())
	randomOne := rand.New(sourceOne)

	fmt.Print(randomOne.Intn(100), ",")
	fmt.Print(randomOne.Intn(100))
	fmt.Println()

	sourceTwo := rand.NewSource(42)
	randomTwo := rand.New(sourceTwo)

	fmt.Print(randomTwo.Intn(100), ",")
	fmt.Print(randomTwo.Intn(100))
	fmt.Println()

	sourceThree := rand.NewSource(42)
	randomThree := rand.New(sourceThree)

	fmt.Print(randomThree.Intn(100), ",")
	fmt.Print(randomThree.Intn(100))
	fmt.Println()
}
