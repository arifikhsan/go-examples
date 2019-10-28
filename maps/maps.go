package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["keyOne"] = 7
	m["keyTwo"] = 13

	fmt.Println("map: ", m)

	v1 := m["keyOne"]
	fmt.Println("variable one: ", v1)

	fmt.Println("length: ", len(m))

	delete(m, "keyTwo")
	fmt.Println("map: ", m)

	_, prs := m["keyTwo"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
}
