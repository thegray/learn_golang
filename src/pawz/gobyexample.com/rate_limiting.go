package main

import "time"
import "fmt"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			fmt.Println("goroutine in action", t)
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 10; i++ {
		fmt.Println("attempt to fill bursty", i)
		burstyRequests <- i
		fmt.Println("fill bursty req", i)
	}
	fmt.Println("0")
	close(burstyRequests)
	fmt.Println("1")
	for req := range burstyRequests {
		fmt.Println("2")
		<-burstyLimiter
		fmt.Println("3")
		fmt.Println("request", req, time.Now())
	}
}
