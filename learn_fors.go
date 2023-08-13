package main

import "fmt"

func print_elements(arr [4]int) {
  for i := 0; i < len(arr); i++ {
    fmt.Println("i: ", i, "value: ", arr[i])
  }
}
func print_slice_arr(arr []int) {
  for i := 0; i < len(arr); i++ {
    fmt.Println("i: ", i, "value: ", arr[i])
  }
}

func main() {
  some_arr := [4]int{1,2,3,4}
  print_elements(some_arr)

  some_slice_arr := []int{1,2,3}
  print_slice_arr(some_slice_arr)
}
