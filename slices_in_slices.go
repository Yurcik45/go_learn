package main

import "fmt"

func use_var() {
  var matrix [][]uint8

  row1 := []uint8{1, 2, 3}
  row2 := []uint8{4, 5, 6}
  row3 := []uint8{7, 8, 9}

  matrix = append(matrix, row1)
  matrix = append(matrix, row2)
  matrix = append(matrix, row3)

  fmt.Println(matrix)
}

func not_use_var() {
  matrix := [][]uint8{
    {1,2,3},
    {4,5,6},
    {7,8,9},
  }
  fmt.Println(matrix)
}

func main() {
  use_var()
  not_use_var()
}

