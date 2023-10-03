// +build !excludeUnused
package main

import (
  "fmt"
  "io"
  "strings"
)

func main() {
  eof_msg := io.EOF
  buff := make([]byte, 8)
  r := strings.NewReader("text with spaces")
  fmt.Println(eof_msg, buff)
  fmt.Println("r: ", r)
}

