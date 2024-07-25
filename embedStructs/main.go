/* main.go */

package main

import "fmt"

type Coords struct {
	x, y int
}

type Circle struct {
	radius int
	Coords
}

func (c Circle) getDiameter() int {
	return c.radius * 2
}

func main() {

	var ring Circle

	ring.radius = 15
	ring.x = ring.radius
	ring.y = ring.radius

	fmt.Printf("Diameter: %v\n", ring.getDiameter())
	fmt.Printf("Point X: %v Y: %v \n", ring.x, ring.y)
}
