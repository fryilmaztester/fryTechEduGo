package main

import "fmt"

func main() {

	//naming rules
	//1 -> İlk karakter harf olmalı
	//2--> Camel case kullanılır
	//export olacaksa (yani paket dışında) İlk harf büyük olmalıdır.

	//func: Belirlir bir işlemi ve görevi yapmak icin tanımlanmıs kod blokları

	var x, y, sum int

	x = 5
	y = 10
	sum = x + y

	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum)

	x = 7
	y = 11
	sum = x + y
	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum)

	// Fonk yardımı ile programı daha moduler bir hale getiririz
	// More Readable code
	// From complex to simple

	fmt.Println(sumOf(7, 10)) //17
	myfunc()
	sumTwo(4, 3)
	hello("FRY", 43)
}

func sumOf(num1 int, num2 int) int {

	//func keyword
	// func name
	//()
	//() --> icine eger parametre var ise girilir --> parametre name ve data type
	//return type var ise yazılır
	//{ code Burada return keyword u yazılmalıdır}

	return num1 + num2
}

func sumTwo(num1, num2 int) {
	fmt.Println(num1 + num2)
}

func myfunc() {
	fmt.Println("MyFunc: ", "Hello my func")
}

func hello(name string, age int) {

	fmt.Printf("Adım %s, yaşım %d\n", name, age)
	fmt.Println("name: ", name, " ", "age: ", age)
}
