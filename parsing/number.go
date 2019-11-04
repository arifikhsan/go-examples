package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 32)
	f2, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	fmt.Println(f2)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("798", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}