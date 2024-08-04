/* main.go */

package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	tmpFile, err := os.CreateTemp("", "Data-*")
	check(err)
	fmt.Printf("\nCreated File:\n%v \n", tmpFile.Name())

	nb, err := tmpFile.WriteString("Go Programming Fun!\n")
	check(err)

	txt, err := os.ReadFile(tmpFile.Name())
	check(err)
	fmt.Printf("\nRead: %v bytes - %v \n", nb, string(txt))

	tmpFile.Close()

	os.Remove(tmpFile.Name())

	_, err = os.Stat(tmpFile.Name())
	check(err)
}
