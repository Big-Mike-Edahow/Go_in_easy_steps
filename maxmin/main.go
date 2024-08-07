/* main.go */

package main

import (
	"fmt"
	"math"
)

func main() {
	square := math.Pow(5, 2) // 5 to power 2 (5x5).
	cube := math.Pow(4, 3)   // 4 to power 3 (4x4x4).

	fmt.Println("Largest Positive:", math.Max(square, cube))
	fmt.Println("Smallest Positive:", math.Min(square, cube))

	square *= -1
	cube *= -1

	fmt.Println("\nLargest Negative:", math.Max(square, cube))
	fmt.Println("Smallest Negative:", math.Min(square, cube))
}
