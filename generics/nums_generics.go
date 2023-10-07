package main

import "fmt"

type Number interface {
  int64 | float64
}

func sumInts(ints []int64) int64 {
  var sum int64
  for _, v := range ints {
    sum += v
  }
  return sum
}

func sumFloats(floats []float64) float64 {
  var sum float64
  for _, v := range floats {
    sum += v
  }
  return sum
}

func sumIntsOrFloats1[T int64 | float64](nums []T) T {
  var sum T
  for _, v := range nums {
    sum += v
  }
  return sum
}

func sumIntsOrFloats2[T Number](nums []T) T {
  var sum T
  for _, v := range nums {
    sum += v
  }
  return sum
}

func main() {
  ints := []int64{1,2,3,4,5}
  floats := []float64{6.2,7.4,8.9}

  fmt.Println("common sum ints: ", sumInts(ints))
  fmt.Println("common sum floats: ", sumFloats(floats))

  fmt.Println("gen var1 sum ints: ", sumIntsOrFloats1(ints))
  fmt.Println("gen var1 sum floats: ", sumIntsOrFloats1(floats))

  fmt.Println("gen var2 sum ints: ", sumIntsOrFloats2(ints))
  fmt.Println("gen var2 sum floats: ", sumIntsOrFloats2(floats))
}
