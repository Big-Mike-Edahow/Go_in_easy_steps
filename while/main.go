// main.go

package main

import "fmt"

func main() {

	var counter int = 1

	for counter <= 5 {
		fmt.Println("Loop Iteration", counter)
		counter++
	}

	var i int = 10
	for {
		fmt.Println("\t\t\tCountdown", i)
		i--
		if i < 1 {
			fmt.Println("\t\t\t\t\tLift Off!")
			break
		}
	}

}
