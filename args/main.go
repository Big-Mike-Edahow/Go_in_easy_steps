// main.go

package main

import "fmt"

func main() {

	num := 5
	fmt.Println("Original:", num)

	square(num)
	fmt.Println("Original:", num)
}

func square(num int) {

	fmt.Println("\t\tReceived Copy:", num)

	num = num * num
	fmt.Println("\t\tModified Copy:", num)
}
