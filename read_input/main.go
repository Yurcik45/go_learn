package main

import "fmt"

type UserData struct {
  name string
  email string
  age int8
}

func main() {
  var user = UserData{}
  ReadString("name", &user.name)
  ReadString("email", &user.email)
 
  fmt.Println("user: ", user)
}
