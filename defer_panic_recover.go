package main

import "fmt"

func may_panic() {
  panic("controlled panic, all is ok :)")
}

func may_error() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered. Error:\n", r)
    }
  }()
  
  may_panic()
}

func main() {
  fmt.Println("starging")
  may_error()
  fmt.Println("msg after recover")
}
