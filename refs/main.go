// main.go

package main

import "fmt"

func main() {

	var num int = 5
	
	fmt.Println("Original:", num)

	square(&num)
	fmt.Println("Original:", num)
}

func square(num *int) {

	fmt.Println("\t\tReceived Address:", num)

	*num = *num * *num
	fmt.Println("\t\tModified Original:", *num)
}
