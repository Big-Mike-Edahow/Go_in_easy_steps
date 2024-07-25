/* main.go */

package main

import "fmt"

func main() {

	var num int = 20
	var ptr *int = &num

	fmt.Printf("num value: %v \t\t\ttype: %T \n", num, num)
	fmt.Printf("num address: %v \ttype: %T \n\n", ptr, ptr)

	fmt.Printf("num via pointer: %v \t\ttype: %T \n", *ptr, *ptr)
	fmt.Printf("ptr address: %v \ttype: %T\n\n", &ptr, &ptr)

	*ptr = 100
	fmt.Printf("new num value: %v \t\ttype: %T \n", num, num)

}
