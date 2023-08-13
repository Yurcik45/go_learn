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

func print_obj_pairs(obj map[string]interface{}) {
  for k, v := range obj {
    fmt.Println("key: ", k, "value: ", v)
  }
}

type User struct {
  name string
  age int
}

func main() {
  some_arr := [4]int{1,2,3,4}
  print_elements(some_arr)

  some_slice_arr := []int{1,2,3}
  print_slice_arr(some_slice_arr)

  some_obj := map[string]interface{}{
    "name": "Rusik",
    "age": 90,
  }
  print_obj_pairs(some_obj)

  for _, c := range "abcd" {
    fmt.Print(" ", c)
  }
  fmt.Println()

  user := User{name: "yurcik", age: 20}
  fmt.Println("user str: ", user.name)
}
