package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for index, num := range nums {
		if num == 3 {
			fmt.Println("index: ", index)
		}
	}

	keyValues := map[string]string{
		"a": "apple",
		"b": "banana"}

	for key, value := range keyValues {
		fmt.Println(key, "->", value)
	}

	for key := range keyValues {
		fmt.Println("key:", key)
	}

	for index, unicode := range "go" {
		fmt.Println(index, unicode)
	}
}
