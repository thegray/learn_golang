package main

import (
	"fmt"
	"os"
)

var c = fmt.Printf

type point struct {
	x, y int
}

func (gg *point) printPoint() {
	c("%#v\n", gg)
}

func main() {
	p := point{1, 2}

	c("%v\n", p)
	c("%+v\n", p)
	c("%#v\n", p)
	c("%T\n", p)
	c("%t\n", true)
	c("%d\n", 123)
	c("%b\n", 14)
	c("%c\n", 33)
	c("%x\n", 456)
	c("%f\n", 78.9)
	c("%e\n", 123400000.0)
	c("%E\n", 123400000.0)
	c("%s\n", "\"string\"")
	c("%q\n", "\"string\"")
	c("%x\n", "hex this")
	c("%p\n", &p)
	c("|%6d|%6d|\n", 12, 345)
	c("|%6.2f|%6.2f|\n", 1.2, 3.45)
	c("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	c("|%6s|%6s|\n", "foo", "b")
	c("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

	p.printPoint()
}
