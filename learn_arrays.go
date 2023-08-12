package main

import "fmt"

func print_space() {
  fmt.Println()
}

func test_param(arr [2]int) {
  fmt.Println("test arr: ", arr)
}

func test_param_ptr(arr *[2]int) {
  fmt.Println("test ptr: ", arr)
  arr[0] = 22
}

func main() {
  // with const length
  const_len_arr := [2]int{1,2}
  fmt.Println("with const len: ", const_len_arr)
  test_param(const_len_arr)
  test_param_ptr(&const_len_arr)
  fmt.Println("arr at the end: ", const_len_arr)
  print_space()

  // just initialization
  var colors [2]string
  colors[0] = "white"
  colors[1] = "black"
  fmt.Println("just init: ", colors)
}
