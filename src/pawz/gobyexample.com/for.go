package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without condition will lopp forever until break out or return from the function
	for {
		fmt.Println("loop!")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			fmt.Println("genap!")
			continue
		}
		fmt.Println(n)
	}
}
