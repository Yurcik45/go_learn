package main

import (
  "fmt"
)

func indexOfStrings(arr *[]string, value *string) int {
  index := -1
  for i := 0; i < len(*arr); i++ {
    if (*arr)[i] == *value {
      index = i
    }
  }
  return index
}

func indexOfAny [T comparable](arr []T, value T) int {
  index := -1
  for i := 0; i < len(arr); i++ {
    if arr[i] == value {
      index = i
      break
    }
  }
  return index
}


func tests() {
  cars := []string{ "audi", "mercedes", "bmw" }
  to_find := cars[1]
  fmt.Println("cars: ", cars)
  fmt.Println("to find: ", to_find)
  
  method_1_index := indexOfStrings(&cars, &to_find)
  fmt.Println("method 1 index: ", method_1_index)
}

func main() {
  tests()
  cars := []string{ "audi", "mercedes", "bmw" }
  numbers := []int{ 1, 22, 53, 75, 84, 34 }
  floats := []float32{ 1.3, 4.5, 7.2, 22.5 }


  test1 := indexOfAny(cars, "bmw")
  test2 := indexOfAny(numbers, 84)
  test3 := indexOfAny(floats, 4.5)

  fmt.Println("\n", test1, "\n", test2, "\n", test3)
}
