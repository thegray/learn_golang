package main

import "fmt"

type rect struct {
	width, height int
}

type rekt struct {
	gg, wp int
}

//what the fuck with this method receiver
func (r *rect) area() int {
	return r.width * r.height
}

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
