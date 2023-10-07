package main

import "fmt"
import "time"

func test_pr[T string | int](text T) {
  fmt.Println(text)
}

func main() {
  go test_pr(1)
  go test_pr(2)
  go test_pr(3)
  go test_pr(4)
  go test_pr(5)

  time.Sleep(time.Millisecond)
}
