/* main.go */

package main

import "fmt"

type Bird interface {
	speak() string
	move() string
}
type Parrot struct{}

func (Parrot) speak() string {
	return "Squawk, squawk!"
}
func (Parrot) move() string {
	return "Scully says, \"A parrot flies away.\""
}

type Chicken struct{}

func (Chicken) speak() string {
	return "Cluck, cluck!"
}
func (Chicken) move() string {
	return "Foghorn says, \"Chickens cannot fly.\""
}

func nudge(b Bird) {
	fmt.Printf("%v \n", b.speak())
	fmt.Printf("%v \n", b.move())
}

func main() {
	var skully Parrot
	var foghorn Chicken

	nudge(skully)
	fmt.Println()
	nudge(foghorn)
}
