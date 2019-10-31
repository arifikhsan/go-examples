package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty slice: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	l := s[2:5]
	fmt.Println("slice [2:5]: ", l)

	l = s[:5]
	fmt.Println("slice up to exclude 5: ", l)

	l = s[2:]
	fmt.Println("slice start from 2: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declare and initialize: ", t)

	twoDimensional := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoDimensional[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoDimensional[i][j] = i + j
		}
	}
	fmt.Println("twoDimensional: ", twoDimensional)
}
