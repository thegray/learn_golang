package main

import "fmt"
import "math"

// basic interface for geometric shapes
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

// To implement an interface in Go, we just need to implement all the methods in the interface.
// Here we implement geometry on rects.
func (r rect) area() float64 {
	return r.w * r.h
}

func (r rect) perim() float64 {
	return 2*r.w + 2*r.h
}

// The implementation for circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call methods that are in the named interface.
// Here’s a generic measure function taking advantage of this to work on any geometry.
func measure(g geometry) {
	fmt.Println("the geometry :", g)
	fmt.Println("the area :", g.area())
	fmt.Println("the perim :", g.perim())
}

func main() {
	r := rect{w: 3, h: 4}
	c := circle{radius: 5}

	// The circle and rect struct types both implement the geometry interface so
	// we can use instances of these structs as arguments to measure.
	measure(r)
	measure(c)

}
