/* main.go */

package main

import "fmt"

type Member struct {
	firstName string
	lastName  string
}

func (m Member) fullName() string {
	return m.firstName + " " + m.lastName
}

type Article struct {
	title string
	body  string
	Member
}

func (a Article) content() {
	fmt.Println("Title:", a.title)
	fmt.Println("Content:", a.body)
	fmt.Printf("Author: %v \n\n", a.fullName())
}

func main() {

	member1 := Member{
		"Mike",
		"McGrath",
	}

	article1 := Article{
		"Object Oriented Programming",
		"In Go, Composition emulates Inheritance",
		member1,
	}
	article1.content()

	article2 := Article{
		"Object Oriented Programming",
		"Coming next... Polymorphism",
		member1,
	}
	article2.content()
}
