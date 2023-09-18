package main

import (
	"fmt"
)

type Colo interface {
	Cleaner()
}

type Circle struct {
	Radius float64
}

type Octaider struct {
	Radius float64
}

func (c *Circle) Cleaner() {
	c.Radius = 0
}

func (some_octaider *Octaider) Cleaner() {
	some_octaider.Radius = 0
}

func printAll(some_circle []Colo) {
	for i := 0; i < len(some_circle); i++ {
		fmt.Println(some_circle[i])
	}
}

func cleanAll(some_circle []Colo) {
	for i := 0; i < len(some_circle); i++ {
		some_circle[i].Cleaner()
	}
}
func main() {
	circle := Circle{Radius: 3.0} // Use a pointer to Circle
	circle_arr := []Colo{&circle, &circle}
	printAll(circle_arr)
	cleanAll(circle_arr)
	printAll(circle_arr)
	fmt.Println("=======")
	// octaider
	octaider := Octaider{Radius: 3.0} // Use a pointer to Circle
	octaider_arr := []Colo{&octaider, &octaider}
	printAll(octaider_arr)
	cleanAll(octaider_arr)
	printAll(octaider_arr)
}
