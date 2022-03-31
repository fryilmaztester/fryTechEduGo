package main

import "fmt"

func main() {

	x, y := 15, 10

	fmt.Printf("%T, %v\n", x, x) //int 10
	fmt.Printf("%T, %v\n", y, y) // int 15

	fmt.Println("x+y: ", x+y)

	fmt.Println("x-y: ", x-y)

	fmt.Println("x*y: ", x*y)

	fmt.Println("x/y: ", x/y)
	//1 --> Her iki değer de int değer olduğu için Go bunun sonucu da int olarak kabul eder ve 1 olarak kabul eder.

	fmt.Println("x%y: ", x%y) //Reminde operator

	//INCREMENT

	f := 8

	fmt.Println("f: ", f) //8

	//1.Method
	f = f + 1

	fmt.Println("Increment Got: ", f) //9

	//2.Method

	f++
	fmt.Println("Increment Got Second: ", f) //10

	//DECREMENT

	e := 6

	fmt.Println("e: ", e) //6

	//1.Method
	e = e - 1

	fmt.Println("Decrement Got: ", e) //5

	//2.Method

	e--
	fmt.Println("Decrement Got Second: ", e) //4

}
