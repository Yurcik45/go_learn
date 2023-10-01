package main

import "fmt"

func main() {
  var sources [6]string
  sources[0] = "https://golang.org/doc/"
  sources[1] = "https://tour.golang.org/welcome/1"
  sources[2] = "https://golang.org/doc/effective_go.html"
  sources[3] = "https://gobyexample.com/"
  sources[4] = "https://play.golang.org/"
  sources[5] = "https://blog.golang.org/"

  for i := 0; i < len(sources); i++ {
    fmt.Print(i + 1, ":\t")
    fmt.Println(sources[i])
  }
}
