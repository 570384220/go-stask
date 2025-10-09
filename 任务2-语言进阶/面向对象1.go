package main

import (
	"fmt"
	"math"
)

func stu19() {
	r := Rectangle{X: 2.0, Y: 3.0}
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())

	c := Circle{Radius: 4}
	//fmt.Println(c.Area())
	fmt.Printf("%.2f\n", c.Area())
	//fmt.Println(c.Perimeter())
	fmt.Printf("%.2f\n", c.Perimeter())
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	X float64
	Y float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.X + r.Y) * 2
}

func (r Rectangle) Area() float64 {
	return r.X * r.Y
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
