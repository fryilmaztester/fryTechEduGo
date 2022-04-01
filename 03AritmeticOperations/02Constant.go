package main

import (
	"fmt"
	"math"
)

//Bir programın çalışma süresi boyunca değişmeyen veri değerleri bölümüdür
//Constant == > Sabit

//Örn: Pi sayısı

func main() {

	r := 3.0

	const pi float64 = 3.14
	areOfCircle := pi * (math.Pow(r, 2))

	fmt.Println("areOfCircle: ", areOfCircle)

	//Constant lar Go da compile time a aittir. const --> Compile Time
	//Değişkenler ise Run time a aittir.        var -----> Var Time

	var x = math.Pow(3, 4)
	//const y = math.Pow(3, 4) //Initialize edemedği için Burada hata verir
	//const y --> bir sabit oluşturudysak go bizden onun değerini ister

	fmt.Println(x)

	f := 3
	const d = 3

	fmt.Printf("%T, %v\n", f, f)
	fmt.Printf("%T, %v\n", d, d)

	u := 5
	const g = 4
	fmt.Println(g + u)

	const e, y = 3, 5

}
