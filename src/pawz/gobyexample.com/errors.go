package main

import "errors"
import "fmt"

// By convention, errors are the last return value and have type error, a built-in interface.
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message.
		return -1, errors.New("cant work with 42")
	}
	// A nil value in the error position indicates that there was no error.
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cant work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 15, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed :", e)
		} else {
			fmt.Println("f1 worked :", r)
		}
	}
	for _, i := range []int{7, 16, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed :", e)
		} else {
			fmt.Println("f2 worked :", r)
		}
	}

	// e is struct of argError
	// check if e string is not null, if not null the error has been thrown
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
		fmt.Println("ok :", ok)
	} else {
		fmt.Println("ok :", ok)
	}

	_, e2 := f2(15)
	// e2 is struct of argError
	// check if e2 string is not null, if not null the error has been thrown
	if ae, ok := e2.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
		fmt.Println("what is ok :", ok)
	} else {
		fmt.Println("what is ok :", ok)
	}
}
