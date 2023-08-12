package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Print
	fmt.Print("This is ", "Print.", " No newline here.")
	fmt.Print(" Another Print.")

	// fmt.Println
	fmt.Println("\nThis is", "Println.")
	fmt.Println("This prints", "with a newline.")

	// fmt.Fprint
	file, _ := os.Create("output.txt")
	defer file.Close()
	fmt.Fprint(file, "This is Fprint.\n")

	// fmt.Fprintln
	fmt.Fprintln(file, "This is Fprintln.")
	fmt.Fprintln(file, "This prints with a newline.")

	// fmt.Errorf
	err := fmt.Errorf("An error occurred: %s", "File not found")
	fmt.Println(err)

	// fmt.Scan
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	fmt.Printf("You entered: %d\n", num)

	// fmt.Sscanf
	var input string = "42 3.14"
	var i int
	var f float64
	n, _ := fmt.Sscanf(input, "%d %f", &i, &f)
	fmt.Printf("Sscanf read %d values: i = %d, f = %f\n", n, i, f)
}
