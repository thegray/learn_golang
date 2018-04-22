package main

import "os"
import "fmt"

// Command-line arguments are a common way to parameterize execution of programs.
// For example, go run hello.go uses run and hello.go arguments to the go program

func main() {

	// os.Args provides access to raw command-line arguments.
	// Note that the first value in this slice is the PATH to the program, and os.Args[1:] holds the actual arguments to the program.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

}
