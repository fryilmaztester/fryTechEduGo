package main

import "fmt"

//IMPORTANT NOTE: Short Declaratıon method must not use on packega scope
var packVar = "Package Scope"

func main() {

	//Scope: Kapsam kavramdır.

	//Her bir değişkenin belirli bir scope u yani kapsama alanı vardır.
	//O değişkene ulaşbileceğimzi bir alan olarka tanımlayabilirriz

	//Package değişkeni
	//Func değişkeni olarak ikiye ayırabiliriz

	funcVar := "Func Var"
	fmt.Println("func scope unu kapsayan bir variabledir. ", funcVar)

	fmt.Println("package duzeyinde tanımlanmış bir varibledir. ", packVar)

	if true {
		var blockScope = "Block Scope"
		fmt.Println("Blok yapısı için tanımlanmıs bir değişkendir. ", blockScope)

		//Bu blok icinde package olarak tanımlanmayan bir değişkene ulaşılamaz.
		//Bu bloktan herhangib bir değişkene ulaşmak istiyorsam ya bu blokta tanımlanmış olmalı ya da
		//package duzeyinde tanımlanmış olması gerekiyor
	}

	myFunc()
}
func myFunc() {

	fmt.Println("package duzeyinde tanımlanmış bir varibledir. ", packVar)
}
