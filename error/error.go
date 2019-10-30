package main

import (
	"errors"
	"fmt"
)

func functionOne(argument int) (int, error) {
	if argument == 42 {
		return -1, errors.New("can't work with 42")
	}
	return argument + 3, nil
}

type argumentError struct {
	argument int
	prob     string
}

func (e *argumentError) Error() string {
	return fmt.Sprintf("%d - %s", e.argument, e.prob)
}

func functionTwo(argument int) (int, error) {
	if argument == 42 {
		return -1, &argumentError{argument, "can't work with it"}
	}
	return argument + 3, nil
}

func main() {
	for _, num := range []int{7, 42} {
		if result, e := functionOne(num); e != nil {
			fmt.Println("function one failed:", e)
		} else {
			fmt.Println("function one worked:", result)
		}
	}

	for _, num := range []int{7, 42} {
		if result, e := functionTwo(num); e != nil {
			fmt.Println("function two failed:", e)
		} else {
			fmt.Println("function two worked:", result)
		}
	}

	_, e := functionTwo(42)
	if result, ok := e.(*argumentError); ok {
		fmt.Println(result.argument)
		fmt.Println(result.prob)
	}
}
