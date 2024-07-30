/* main.go */

package main

import (
	"fmt"
	"math"
)

func main() {
	var rad, area, perim float64

	rad = 4
	fmt.Printf("Radius of Circle: %.2f \n", rad)

	area = math.Pi * (rad * rad)
	fmt.Printf("Area of Circle: %.2f \n", area)

	perim = 2 * (math.Pi * rad)
	fmt.Printf("Perimeter of Circle: %.2f \n", perim)
}
