// main.go

package main

import "fmt"

func main() {

	var zero int = 0
	var num int = 0
	var max int = 1
	var upper byte = 'A'
	var lower byte = 'a'
	var result bool

	result = (num == zero) // 0 == 0.
	fmt.Println("Equality:\tnum == zero\t", result)

	result = (upper == lower) // A == a.
	fmt.Println("Equality:\tupper == lower\t", result)

	result = (max != zero) //1 != 0.
	fmt.Println("Inequality:\tmax != zero\t", result)

	result = (zero > max) // 0 > 1.
	fmt.Println("Greater:\tzero > max\t", result)

	result = (max <= zero) // 1 <= 0.
	fmt.Println("Less or Equal:\tmax <= zero\t", result)

}
