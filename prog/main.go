/* main.go */

package main

import (
	"prog/verify"
	"fmt"
)

func main() {

	for i := 2; i >= -2; i-- {

		res, err := verify.IsPosInt(i)

		if err != nil {
			fmt.Println("Failed:", err)
		} else {
			fmt.Println(res, "Passed!")
		}
	}
}
