package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Println("working..")
	time.Sleep(time.Second * 2)
	fmt.Println("done!")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
