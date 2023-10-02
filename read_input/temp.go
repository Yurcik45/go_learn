package main

import "fmt"

func read_age_safety(age *int8) {
  is_err := true
  for is_err {
    _, err := fmt.Scan(age)
    if err != nil {
      fmt.Println("One more time")
    } else {
      is_err = false
    }
  }
  fmt.Println("age: ", *age)
}

func read_and_validate_age(age *int8) {
  for {
    _, err := fmt.Scan(age)
    if err != nil || *age < 5 {
      fmt.Println("One more time")
    } else {
      break
    }
  }
}

