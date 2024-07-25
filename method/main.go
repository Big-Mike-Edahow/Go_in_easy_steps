package main

import (
	"fmt"
	"method/carData"
)

func main() {

	var porsche = carData.Car{Color: "blue", Body: "coupe"}
	bentley := carData.Car{Color: "green", Body: "saloon"}

	fmt.Println("Porsche paint is", porsche.Color)
	fmt.Println("Porsche style is", porsche.Body)
	fmt.Printf("Porsche is %v\n\n", porsche.Accelerate())

	fmt.Println("Bentley paint is", bentley.Color)
	fmt.Println("Bentley style is", bentley.Body)
	fmt.Print("Bentley is ", bentley.Zooming(), "\n")
}
