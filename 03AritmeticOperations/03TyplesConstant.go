package main

import "fmt"

func main() {

	const x = 16 //Go 'x' sabitinin veri tipini belirli olarak algılamaz. Varsayılan bir veri tipi olarak atama yapar

	fmt.Printf("%T, %v\n", x, x) //int 16

	const name = "FRY"

	fmt.Printf("%T, %v\n", name, name) //string, FRY

}
