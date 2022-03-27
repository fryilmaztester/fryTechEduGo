package main

import (
	"fmt"
	"strconv"
)

func main() {

	// int, floata donusum

	x := 75
	var y float64
	y = float64(x)

	fmt.Println(y)

	//multiple assign

	v := 43
	b := 5

	v, b = b, v

	fmt.Println("v: ", v)
	fmt.Println("b: ", b)

	// non English variable names

	yaş := 27 //Go UTF 8 uyumlu bir dildir.
	fmt.Println("yaş: ", yaş)

	имя := 53
	fmt.Println("имя: ", имя)

	//Shadowing

	n := 5

	if true {

		n := 10 // n yi tekrar dan oluştruduk. Fakat tekrar oluşturmayıp assign işlemi (n = 10) yapsaydık aşağıda bulunan n de buraki değer işle aynı çıkardı
		n++
		fmt.Println("n: ", n) //11

	}

	fmt.Println("n: ", n) //5

	//40 as String

	num := 65
	str := string(num)

	fmt.Println(str) //65 e denk gelen ASCII de ki char ifadeyi gçsterir

	num1 := 65
	str01 := strconv.Itoa(num1)
	fmt.Println(str01) //str olarak yazdıırlmıştır.

}
