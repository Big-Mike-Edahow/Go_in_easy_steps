/* main.go */

package main

import (
	"encap/cube"
	"fmt"
)

func main() {

	var box cube.Dimensions

	box.SetSize(2, 4, 6)

	fmt.Print("Box Dimensions: ", 2, "ft x ", 4, "ft x ", 6, "ft.\n" )
	fmt.Printf("Footprint: %vft.\n", box.GetArea())
	fmt.Printf("Volume: %vft.\n", box.GetVolume())
}
