package main

import "fmt"

type employee struct {
	name      string
	age       int
	isMarried bool
}
type manager struct {
	employee
	hasDegree bool
}

func main() {

	var e1 = employee{
		"Fatih",
		21,
		false,
	}

	fmt.Println(e1)
	e2 := e1
	fmt.Println(e2)

	m1 := manager{

		employee: employee{
			"Ramazan",
			32,
			false,
		},
		hasDegree: true,
	}

	fmt.Println(m1)

	//Anonym Strcy

	theBoss := struct {
		name  string
		money bool
	}{name: "The Boss", money: true}

	fmt.Println(theBoss)

	fmt.Println("Fatih Branch Deneme")
}
