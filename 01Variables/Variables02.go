package main

//Go know that variables type before not execute so Go is fatly programming lang other than

import "fmt"

func main() {

	var name = "Fry"
	fmt.Printf("%T\n", name) // Printf("%T", name) --> it gives us name's data type
	//Go can understand,  assigned value's data type.

	/*
			BASIC DATA TYPES

			-Numeric Types

			uint --> Only positive integer
			int --> Only integer number include negative number
			float --> Only doube-le type
			complex --> the set of complex number, it is for a huge number
		    byte --> unit8
			rune --> int32


		-STRING Data Type

		string --> it is using for string sentences. You must not use negative

		--Boolean Type

		bool --> it is using for loginc sentences it returns true or false

	*/

	/*
		In Go, values have a data type. The data type determines what type of information is being stored and
		how much space is needed to store it. Go has basic data types such as:

		string
		bool

		numeric types:
		int8, uint8, int16, uint16, int32 , uint32, int64, uint64, int, uint, uintptr
		float32, float64
		complex64, complex128
	*/

	var age uint = 9
	fmt.Println(age)

	var street = "Istanbul"
	fmt.Println(street)

	var isRetired = true
	fmt.Println(isRetired)

	//	street = 23 --> You can not use number type for str. Because go knows that street is string data.

}
