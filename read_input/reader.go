package main

import "fmt"

func more() { fmt.Println("One more time") }

func ReadString(name string, s *string) {
  for {
    fmt.Print(name, ": ")
    _, err := fmt.Scan(s)
    switch {
      case err != nil: 
        more()
        fallthrough
      case name == "name" && !ValidateName(*s):
        more()
        fallthrough
      case name == "email" && !ValidateEmail(*s):
        more()
      default: return
    }
  }
}

func ReadInt(name string, i *int) {
  for {
    fmt.Print(name, ": ")
    _, err := fmt.Scan(i)
    if err != nil || !ValidateAge(*i) {
      more()
    } else {
      break
    }
  }
}
