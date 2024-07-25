// main.go

package main

import "fmt"

func factorial(num int) int {

	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	for i := 1; i <= 7; i++ {
		fmt.Println("Factorial", i, "=", factorial(i))
	}
}

