package main

import "fmt"

func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable,
	// e.g. for var f string = "short" in this case.
	f := "short"
	fmt.Println(f)

	g := 9
	fmt.Println(g)

	var gg string
	fmt.Println("var gg = " + gg)

	var aa int
	fmt.Println(aa)

	tes = 500000000 //this will yield compile error

}
