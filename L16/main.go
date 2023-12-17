package main

import (
	"fmt"
	"math"
)

// shape interface
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

//shape info

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 6.4},
		circle{radius: 12.3},
		square{length: 10.4},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}
