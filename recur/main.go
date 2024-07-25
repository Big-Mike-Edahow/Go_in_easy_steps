// main.go

package main

import "fmt"

func main() {
	fmt.Println("Countdown commencing, fire one.")
	countDown(10)
}

func countDown(num int) {
	if num < 1 {
		fmt.Println("\nLift Off!")
	} else {
		fmt.Printf("%d\n", num)
		num--
		countDown(num)
	}
}

