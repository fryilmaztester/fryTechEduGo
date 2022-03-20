package main

import "fmt"

func main() {

	// We use var keyword for create a variable but it is not recommended because
	// var is repeatedly so Go offer us a solution.
	//Our main idea is Don't Repeat Yourself

	/*
		var (
			name       string = "Fatih Ramazan YILMAZ"
			name02            = "FRY"
			name03            = "Ozan"
			schoolName        = "Marmara University"
			isMarried         = true
			age               = 35
		)

	*/

	//Multiple Variable Creation
	/*var name, schoolName, isMarried, age, weight = "fatih", "Marmara", true, 27, 82.5 */

	name, schoolName, isMarried, age, weight := "fatih", "Marmara", true, 27, 82.5

	/*
		name03 := "FRYTechEducation"
		fmt.Println(name03)
	*/
	fmt.Println(weight)
	//fmt.Println(name02)
	fmt.Println(name)
	fmt.Println(isMarried)
	fmt.Println(schoolName)
	fmt.Println(age)

	/*
		In Go, when a variable is declared without initializing a value, it has a default value.
		The default value is known as the zero value.

		Different zero values exist for different data types:

		Type    Zero Value
		ints     0
		floats   0
		string   "" (empty string)
		boolean  false
	*/

	var pcMarker string
	var model int
	var inc float32
	var isSSD bool

	fmt.Println("Not initializing for STRING: ", pcMarker)
	fmt.Println("Not initializing for int: ", model)
	fmt.Println("Not initializing for float:", inc)
	fmt.Println("Not initializng for boolean: ", isSSD)

}
