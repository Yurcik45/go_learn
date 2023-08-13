package main

import (
  "fmt"
)

type User struct {
  name string
  age int
}

func print_space() {
  fmt.Println("--------------------------------------------------------------------")
}

func main() {
  // use structure
  user_st := User{ name: "somebody", age: 22 }
  fmt.Println("user use struct: ", user_st)
  print_space()

  // use pointer to structure
  user_ptr := &User{ name: "whoami", age: 42 }
  fmt.Println("user_prt: ", user_ptr)
  fmt.Println("*user_ptr: ", *user_ptr)
  print_space()

  // use map with strong types
  some_params_f := map[string]string{ "paramf1": "valuef1", "paramf2": "valuef2" }
  fmt.Println("params use map: ", some_params_f)
  print_space()

  // use map with interface
  user_map := map[string]interface{}{
    "name": "somebody",
    "age": 22,
  }
  fmt.Println("user use map: ", user_map)
  print_space()

  // use map with struch
  user_mst := map[int]User{ 0: User{ name: "somebody2", age: 44 } }
  fmt.Println("user_mst use map with struct: ", user_mst)
  fmt.Println("user_mst[0]use map with struct: ", user_mst[0])
  print_space()

  // use make with strong types
  colors := make(map[string]string)
  colors["black"] = "#000"
  colors["white"] = "#fff"
  fmt.Println("colors use make: ", colors)
  print_space()

  // use make with interface
  some_params_s := make(map[string]interface{})
  some_params_s["params1"] = "values1"
  some_params_s["params2"] = 2
  fmt.Println("params use make with interface: ", some_params_s)
  print_space()

  // use make with struct
  users := make(map[int]User)
  users[0] = User{ name: "yurcik", age: 24 }
  fmt.Println("use maps users: ", users)
  fmt.Println("use maps Users[0]: ", users[0])
  print_space()

  {
    // create an empty map without make function
    some_params_f := map[string]string{}
    fmt.Println("params use map: ", some_params_f)
    print_space()
  }

  {
    // create an empty map with make function
    map_blyat := make(map[string]string)
    fmt.Println("params use map: ", map_blyat)
    print_space()
  }
}
