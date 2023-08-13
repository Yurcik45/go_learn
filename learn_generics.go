package main

import (
  "fmt"
)

func Get_map_keys[K comparable, V any] (obj map[K]V) []K {
  keys := make([]K, 0, len(obj))
  for key:= range obj {
    keys = append(keys, key)
  }
  return keys
}

func main() {
  some_obj := make(map[string]string)
  fmt.Println("initial obj: ", some_obj)
  some_obj["name"] = "yurcik"
  some_obj["marry"] = "kakish"

  keys := Get_map_keys[string, string] (some_obj)
  fmt.Println("keys in result: ", keys)
}
