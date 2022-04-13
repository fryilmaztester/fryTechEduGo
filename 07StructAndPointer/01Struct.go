package main

import "fmt"

func main() {

	/*var employee struct {
		name      string
		age       int
		isMarried bool
	}

	fmt.Println(employee) //{ 0 false}

	employee.name = "Gurcan"
	employee.age = 21
	employee.isMarried = false

	fmt.Println(employee)*/

	type employee struct {
		name      string
		age       int
		isMarried bool
	}

	var e1 = employee{
		"Fatih",
		21,
		false,
	}

	fmt.Println(e1) //{Fatih 21 false}

}
