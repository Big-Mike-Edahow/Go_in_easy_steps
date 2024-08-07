/* main.go */

package main

import (
	"fmt"
	"strings"
)

func main() {

	s1, s2 := "The Truth is rarely Pure ", "and never Simple."
	str := s1 + s2

	chars, err := fmt.Printf("%v \n", str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bytes Written: ", chars)
		fmt.Println("String Length:", len(str))
	}

	arr := []string{"\n\tStrive", "For", "Greatness!"}
	ast := strings.Join(arr, " ")

	fmt.Println(ast)
	if ast[0] == 10 && ast[1] == 9 {
		fmt.Printf("1st Char: ASCII %v Newline\n", ast[0])
		fmt.Printf("2nd Char: ASCII %v Tab\n", ast[1])
	}
}
