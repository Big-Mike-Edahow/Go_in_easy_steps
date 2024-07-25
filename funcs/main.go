// main.go

package main

import "fmt"

func first() {
	msg := "Hello from the 1st function!"
	fmt.Println(msg)
}

func second() {
	var status bool = true
	var intNum int = 10
	var floatNum float64 = 3.14159
	fmt.Printf("Hello from the 2nd function!\n")
	fmt.Printf("Here are some local variables:\n")
	fmt.Printf("A boolean: %v, an integer: %d, and a float: %.5f.\n", status, intNum, floatNum)
}

func sqFive() {
	fmt.Printf("%v \n", 5*5)
}

func main() {

	first()
	fmt.Print("5 x 5 = ")
	sqFive()
	second()
}
