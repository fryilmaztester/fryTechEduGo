package main

import (
	"fmt"
	"math"
)

func main() {

	//Array: Aynı veri tipine ait birden fazla değişken, eleman var ise onları bir liste içerisinde tutmaya çalışırız

	/*	city1 := "İstanbul"
		city2 := "Roma"
		city3 := "Tahran"
		city4 := "Belgrad"

		fmt.Println(city1, city2, city3, city4)
	*/

	//Array Nasıl Oluşturulur

	//1.Method --> Array in eleman sayısını vererek.
	//cities := [4]string{"İstanbul", "Roma", "Tahran", "Belgrad"} //[İstanbul Roma Tahran Belgrad]

	//2.Method --> Array in eleman sayısnını boş bırakarar
	//cities := []string{"İstanbul", "Roma", "Tahran", "Belgrad"} //[İstanbul Roma Tahran Belgrad]

	//3.Method --> Array in eleman saytısını ... ile göstererek
	cities := [...]string{"İstanbul", "Roma", "Tahran", "Belgrad"} //[İstanbul Roma Tahran Belgrad]
	fmt.Println(cities)

	fmt.Println(cities[0]) //İstanbul birinci indexteki elamnı bize verir

	fmt.Println(len(cities)) //4 Array n uzunluğunu bizlere verir

	/*var myArray [5]int

	fmt.Println(myArray)
	myArray[0] = 4
	fmt.Println(myArray)*/

	/*
		var myArray [4]int
		var myArrayTwo [5] int

		if myArray == myArrayTwo {
			fmt.Println("MESAJJJ")
		}*/

	//Invalid operation: myArray == myArrayTwo (mismatched types [4]int and [5]int)

	for i := 0; i < len(cities); i++ {
		fmt.Print(i, " ", cities[i], " ")
	}
	fmt.Println()

	fmt.Println("-----------------------")

	myNumArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	myNumArr = mySquare(myNumArr)
	fmt.Println(myNumArr)
}

func mySquare(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {
		arr[i] = int(math.Pow(float64(arr[i]), 2))
	}
	return arr
}
