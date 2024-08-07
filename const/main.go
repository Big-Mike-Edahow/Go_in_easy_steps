package main

import (
	"fmt"
	"math"
)

func main() {

	const pi = math.Pi

	const (
		red = iota + 1
		yellow
		green
		brown
		blue
		pink
		black
	)

	fmt.Printf("Pi approximately: %v \n\n", pi)

	fmt.Printf("Red: %v point \n", red)
	fmt.Printf("Blue: %v points \n", blue)
	fmt.Printf("Black: %v points \n", black)

}
