/* main.go */

package main

import (
	"fmt"
	"group/empData"
)

func main() {

	var coder empData.Employee
	coder.Id = 001
	coder.Name = "Alice"
	coder.Dept = "I.T."

	clerk := empData.Employee{Id: 002, Name: "Burt", Dept: "Payroll"}

	fmt.Println(coder)
	fmt.Println(clerk)

	fmt.Printf("\n%v works in %v \n", coder.Name, coder.Dept)
	fmt.Printf("\n%v works in %v \n", clerk.Name, clerk.Dept)
}
