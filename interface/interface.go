package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	circumference() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (instance rectangle) area() float64 {
	return instance.width * instance.height
}

func (instance rectangle) circumference() float64 {
	return 2*instance.width + 2*instance.height
}

func (instance circle) area() float64 {
	return math.Pi * instance.radius * instance.radius
}

func (instance circle) circumference() float64 {
	return 2 * math.Pi * instance.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.circumference())
}

func main() {
	rectangle := rectangle{width: 3, height: 4}
	circle := circle{radius: 5}

	measure(rectangle)
	measure(circle)
}
