/* main.go */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	txt, err := ioutil.ReadFile("./Oscar.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(txt))

	file, err := os.Open("./Oscar.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	pos, err := file.Seek(42, 0)
	if err != nil {
		fmt.Println(err)
	}

	slice := make([]byte, 15)
	nb, err := file.Read(slice)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n%v bytes @ %v: ", nb, pos)
	fmt.Printf("%v\n", string(slice[:nb]))
}
