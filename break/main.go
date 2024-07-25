// main.go

package main

import "fmt"

func main() {

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {

			if i == 3 && j == 2 {
				fmt.Printf("*\t\t\tContinues When i=%v and j=%v\n",i, j)
				continue
			}

			if i == 2 && j == 2 {
				fmt.Printf("*\t\t\tBreaks When i=%v and j=%v\n", i, j)
				break
			}

			fmt.Println("Running i=", i, "j=", j)

		}
	}
}
