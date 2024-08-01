/* main.go */

package main

import (
	"flag"
	"fmt"
)

func main() {
	text := flag.String("text", "C#", "A string")
	number := flag.Int("number", 8, "An integer")
	status := flag.Bool("status", true, "A boolean")
	flag.Parse()
	fmt.Println("Text:", *text, "Number:", *number, " Status:", *status)
}
