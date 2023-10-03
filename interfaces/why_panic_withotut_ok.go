package main

import "fmt"

func common() {
  fmt.Println("common \n")
	var i interface{} = "вітаю"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)
}



func test1() {
  fmt.Println("test 1 \n")
	var i interface{} = "вітаю"

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // паніка!
	fmt.Println(f)
}

func test2() {
  fmt.Println("test 2 \n")
	var i interface{} = "вітаю"

	f, ok := i.(float64)
	fmt.Println(f, ok)
}

func main() {
  common()
  test1() // without 'ok', why 'panic'?
  test2() // with 'ok' and no 'panic' ...
}
