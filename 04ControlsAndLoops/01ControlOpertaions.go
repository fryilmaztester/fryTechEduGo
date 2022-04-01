package main

import "fmt"

func main() {

	x, y := 3, 7

	fmt.Printf("%T, %v\n", x == y, x == y) //bool, false

	//Data type ı eşit olmayan veriler karşılaştırılamaz
	//Go da Bir birinin ataybileceğiniz değerleri karşılatırabilirsiniz

	f, r := 15, 7
	set1 := (f == r)
	set2 := (f > r)

	fmt.Printf("%T, %v\n", set1, set1) //bool false
	fmt.Printf("%T, %v\n", set2, set2) //bool true

	fmt.Printf("%T, %v\n", set1 && set2, set1 && set2) //false && true --> bool false
	fmt.Printf("%T, %v\n", set1 || set2, set1 || set2) //false veya true --> bool true

}
