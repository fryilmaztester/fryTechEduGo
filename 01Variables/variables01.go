package main

import "fmt"

func main() {

	/*
		var name string = "Fatih Ramazan"
		fmt.Println("Hello World", name)

	*/

	//How Can I create variable?

	//1) Varible Decleration
	//2) Variable Name
	//3) Variable Type
	//4) Equals
	//5) Variable Assign
	var name string = "FRY"

	fmt.Println("name: ", name)

	var age int = 27
	fmt.Println("age: ", age)

	var isMarried bool = false
	fmt.Println("isMarried: ", isMarried)

	/*
		var number int = "testing"
		fmt.Println(number)

	*/

	//Go Prog Lang has got static type. You can not assigned a string to int type.

	var firstName, lastName string = "Fatih", "Yılmaz"

	fmt.Println(firstName, lastName) //Fatih Yılmaz

	//Shorltly MEthod
	//you can not use "var" keyword

	isRetired := true
	fmt.Println(isRetired)

	//*** You can change variable value other lines.

	//But  you must be careful " MUST USE ":" markdown in case of "var" ."

	n := 37
	fmt.Println("first n: ", n)

	n = 45
	fmt.Println("second n: ", n)
}
