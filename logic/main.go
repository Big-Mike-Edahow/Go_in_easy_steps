// main.go

package main

import "fmt"

func main() {

	var yes, no bool = true, false

	fmt.Printf("yes = true no = false\n\n")
	result := (yes && no)
	fmt.Println("AND Logic:\tyes && no\t", result)
	result = (yes || no)
	fmt.Println("OR  Logic:\tyes || no\t", result)
	result = !yes
	fmt.Println("NOT Logic:\t!no = ", !no, "\t !yes =", result)
}
