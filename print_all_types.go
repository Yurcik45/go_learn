package main

import "fmt"

func printAllTypes() {
    num := 42
    text := "Hello"
    flag := true
    char := 'A'
    floatNum := 3.14
    pointer := &num

    fmt.Printf("Integer: %d, String: %s, Boolean: %t, Character: %c\n", num, text, flag, char)
    fmt.Printf("Floating-point: %f, Pointer: %p\n", floatNum, pointer)
}

func main() {
    printAllTypes()
}
