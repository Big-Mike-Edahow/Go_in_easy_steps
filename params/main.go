/* main.go */

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Printf("Argument %v: %v \n", i, v)
	}
	fmt.Println("Last Argument:", os.Args[len(os.Args)-1])
}
