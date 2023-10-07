package main

import "fmt"
import "time"

func test_pr[T string | int](text T) {
  fmt.Println(text)
}

func main() {
  go func(text string) {
    fmt.Println(text)
  }("test 1")

  go test_pr(1)
  go test_pr(2)
  go test_pr(3)
  go test_pr(4)
  go test_pr("five")
  go test_pr(6)

  go func(text string) {
    fmt.Println(text)
  }("test 2")

  time.Sleep(time.Millisecond)
}
