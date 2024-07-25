// main.go

package main

import "fmt"

func main() {

	array := [3]string{"Dodge", "Ford", "GMC"}

	slice := array[:]

	list(slice...)

	slice = append(slice, "Porsche", "Ferrari", "Lamborghini")

	list(slice...)

}

func list(autos ...string) {

	for i, v := range autos {
		fmt.Printf("\n%v. %v", i, v)
	}
	fmt.Println()
}
