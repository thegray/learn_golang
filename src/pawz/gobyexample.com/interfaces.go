package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	w, h float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.w * r.h
}

func (r rect) perim() float64 {
	return 2*r.w + 2*r.h
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("the geometry :", g)
	fmt.Println("the area :", g.area())
	fmt.Println("the perim :", g.perim())
}

func main() {
	r := rect{w: 3, h: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

}
