package main

import (
  "fmt"
)

type UserData struct {
  name string
  surname string
  age int8
}

func (u UserData) fullName() string {
  return u.name + " " + u.surname
}

func example1() {
  user := UserData{ "yurcik", "45", 24 }
  full := user.fullName()
  fmt.Println("user: ", user)
  fmt.Println("full: ", full)
}

func main() {
  example1()
}
