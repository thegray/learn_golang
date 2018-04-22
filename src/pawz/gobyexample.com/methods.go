package main

import "fmt"

// Go supports methods defined on struct types.

type rect struct {
	width, height int
}

type rekt struct {
	gg, wp int
}

// This area method has a receiver type of *rect.
// So, this type of declaration (c style afaik), will automatically bind the method(area) to the receiver type(struct rect)
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// Here’s an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (g rekt) asuw(i int) int {
	return (g.gg * g.wp) + i
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area  : ", r.area())
	fmt.Println("perim : ", r.perim())

	// Go automatically handles conversion between values and pointers for method calls.
	// You may want to use a pointer receiver type to avoid copying on method calls or
	// to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("area  : ", rp.area())
	fmt.Println("perim : ", rp.perim())

	g := rekt{gg: 10, wp: 2}
	fmt.Println("asuw  : ", g.asuw(5))

	var gasuw = (rekt).asuw
	fmt.Println("gasuw : ", gasuw(g, 1))

	var pasuw = (*rekt).asuw
	fmt.Println("pasuw : ", pasuw(&g, 2))

}
