package main

import "fmt"

func main() {

	/*
		x := 10
		y := 10.0

		fmt.Printf("%v %T", x, x)
		fmt.Println()
		fmt.Printf("%v %T\n", y, y)

		//fmt.Println(x + y) // Farklı bir data type ait olan değişkenleri biribirleri ile işleme alamayoız


	*/

	z := 10.3
	d := 12

	fmt.Println("------------")
	e := z + float64(d)
	fmt.Println(e)

	var x int8 = 127
	var y int16

	//y = x //x in veri tipi ile y nin veri tipi biribirnden farklı olduğu için bir biri ile swap işlemi yapamayız

	y = int16(x)
	fmt.Println(y)

	//Eğer dönüşümü yapmak isteidğimiz veri tipi diğer veri tipinin sınıtrlarını aşıyorsa en küçük değerini alır

	/*f := 10
	e := "10"

	fmt.Println(f + e) // Str bir ifade type conversation ile numeric bir ifadeye dönüştüremeyiz. Farklı bir metot kullanılabilir. */

	num1 := 106
	str := string(num1)

	fmt.Printf("%v %T\n", num1, num1)
	fmt.Printf("%v %T\n", str, str) //str ye donuşum yapılmıştır ve go bunu ascıı table da ki değere göre alığ "j" olarka okumuştur

}
