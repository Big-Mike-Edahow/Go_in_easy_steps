package main

import "fmt"

type coords struct {
	x, y int
}

type circle struct {
	radius int
	coords
}

func (c circle) getDiameter() int {
	return c.radius * 2
}

func main() {

	var ring circle

	ring.radius = 15
	ring.x = ring.radius
	ring.y = ring.radius

	fmt.Printf("Diameter:%v\n", ring.getDiameter())
	fmt.Printf("Point X:%v Y:%v \n", ring.x, ring.y)
}
