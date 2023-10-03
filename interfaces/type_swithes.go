package main

import (
  "fmt"
)

func sw(i interface{}) {
  switch v := i.(type) {
	  case int:
		  fmt.Printf("Integer: %d\n", v)
    case int8:
      fmt.Printf("Integer 8: %d\n", v)
	  case string:
		  fmt.Printf("String: %s\n", v)
	  default:
		  fmt.Printf("Unknown type\n")
	}
}

func main() {
  var in8 int8 = 8
  var in int = 88
  sw(in8)
  sw(in)
  sw("str")
  sw(30.1)
  sw(false)
}
