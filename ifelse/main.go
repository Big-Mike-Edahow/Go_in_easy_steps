// main.go

package main

import "fmt"

func main() {
	var letterOne byte = 'A'
	var letterTwo byte = 'A'

	if letterOne == letterTwo {
		fmt.Println("Characters Match")
	}

	if 5 > 1 {
		if 7 > 2 {
			fmt.Println("\nBoth Expressions are true")
		}
	}

	if 5 < 1 {
		fmt.Println("\n1st Condition is true")
	} else if letterOne != letterTwo {
		fmt.Println("\n2nd Condition is true")
	} else {
		fmt.Println("\nBoth Conditions are false")
	}
}
