/* main.go */

package main

import "fmt"

func main() {
	a := 8
	b := 4

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Println("Addition:\t", (a + b))
	fmt.Println("Subtraction:\t", (a - b))
	fmt.Println("Multiplication:\t", (a * b))
	fmt.Println("Division:\t", (a / b))
	fmt.Println("Remainder:\t", (a % b))

	a++
	b--
	fmt.Println("Increment:\t", a)
	fmt.Println("Decrement:\t", b)
}
