package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var (
        i int
        f float64
        s string
        b bool
    )

    // Size of different variables in bytes
    sizeInt := unsafe.Sizeof(i)
    sizeFloat := unsafe.Sizeof(f)
    sizeString := unsafe.Sizeof(s)
    sizeBool := unsafe.Sizeof(b)

    fmt.Printf("Size of int: %d bytes\n", sizeInt)
    fmt.Printf("Size of float64: %d bytes\n", sizeFloat)
    fmt.Printf("Size of string: %d bytes\n", sizeString)
    fmt.Printf("Size of bool: %d bytes\n", sizeBool)
}

