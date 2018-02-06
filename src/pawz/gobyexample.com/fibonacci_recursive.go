package main

import "fmt"

func fiboRec(n int) int {
  if n < 1 {
    return 0
  }
  if n == 1 {
    return 1
  }
  return fiboRec(n - 1) + fiboRec(n - 2)
}

func main() {
  d := 10
  fmt.Printf("fibonacci ke-%d : %d \n", d, fiboRec(d))
}
