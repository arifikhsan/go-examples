package main

import (
	"fmt"
	"time"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(time.Now()))

	var now time.Time;
	now = time.Now()
	fmt.Println(now)

	fmt.Println(time.Now().Local())
	fmt.Println(time.Now().Location())
	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().MarshalJSON())
	fmt.Println(time.Now().MarshalText())
}
