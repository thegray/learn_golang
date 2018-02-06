package main

import (
	"fmt"
	"os"
)

func apalah(a, b int) (int, bool) {
	if b > 0 {
		return a + b, true
	} else {
		return a - b, false
	}
}

type Imba struct {
	GG string
	WP int
}

func alala(im Imba) {
	fmt.Println("WP imba in function: ", im.WP)
	im.WP += 1000
	fmt.Println("add WP in function: ", im.WP)
}

func cool(im *Imba) {
	fmt.Println("WP imba in function: ", im.WP)
	im.WP += 1000
	fmt.Println("add WP in function: ", im.WP)
}

func notGGlah(im *Imba) {
	im = &Imba{"nope", 0}
}

func GGla(im *Imba) {
	*im = Imba{"Paul Imba", 10000}
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	} else {
		fmt.Println("uh yeah", os.Args[0])
	}
	fmt.Println("It's over", os.Args[1])

	bisa, gini, ya := "emba", 9999, true
	fmt.Println("bisa: ", bisa)
	fmt.Println("gini: ", gini)
	fmt.Println("ya: ", ya)

	_, aa1 := apalah(1, 5)
	fmt.Println("ret bool: ", aa1)

	aa2, _ := apalah(3, 0)
	fmt.Println("ret bool: ", aa2)

	imba1 := Imba{GG: "Paul", WP: 5000}
	imba2 := &Imba{GG: "Paw", WP: 6000}

	fmt.Println()

	fmt.Println("imba1 WP: ", imba1.WP)
	alala(imba1)
	fmt.Println("imba1 WP after function call: ", imba1.WP)

	fmt.Println("imba2 WP: ", imba2.WP)
	cool(imba2)
	fmt.Println("imba2 WP after function call: ", imba2.WP)

	notGGlah(imba2)
	fmt.Println("imba2 wow: ", imba2.GG, imba2.WP)

	GGla(imba2)
	fmt.Println("imba2 new: ", imba2.GG, imba2.WP)

	/////last////////
}
