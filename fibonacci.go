package main

import "fmt"

func fibonacci() func() int {
    a, b := 0, 1 // Initialize the first two Fibonacci numbers

    return func() int {
        result := a
        a, b = b, a+b // Update a and b for the next Fibonacci number
        return result
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}

