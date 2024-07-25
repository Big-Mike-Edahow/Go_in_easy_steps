// main.go

package main

import "fmt"

func main() {

	num := 2
	char := 'C'

	switch num {
	case 1:
		fmt.Println("Number is One")
	case 2:
		fmt.Println("Number is Two")
	case 3:
		fmt.Println("Number is Three")
	default:
		fmt.Println("Number is Unrecognized")
	}

	switch char {
	case 'A':
		fmt.Println("\nLetter is A")
	case 'B':
		fmt.Println("\nLetter is B")
	default:
		fmt.Println("\nLetter is Unrecognized")
	}

}
