/* main.go */

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("lines.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"Welcome to the wonderful world of Go.",
		"Go is a compiled language that is easy to learn.",
		"Go is fun to use and play around with.",
	}

	for _, v := range d {
		fmt.Fprintln(f, v)
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File written successfully.")
}
