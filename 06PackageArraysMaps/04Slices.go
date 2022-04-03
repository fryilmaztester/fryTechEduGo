package main

import "fmt"

func main() {

	myArr := [3]int{1, 2, 3}
	fmt.Println(myArr) //[1 2 3]

	myArr2 := [...]int{4, 5, 6, 7}
	fmt.Println(myArr2)

	fmt.Println(len(myArr2)) //4

	//slices eleman sayısını bilmez

	mySlc := []int{7, 8, 9}
	fmt.Println(mySlc) //[7 8 9]

	var arr [4]int
	fmt.Println(arr) //[0 0 0 0]
	arr[2] = 4
	fmt.Println(arr) //[0 0 4 0]

	var slc []int
	fmt.Println(slc) //[]
	//slc[0] = 10
	//fmt.Println(slc) //panic: runtime error: index out of range [0] with length 0

	var mySlc01 []int
	mySlc01 = make([]int, 4)
	mySlc01[0] = 10
	fmt.Println(mySlc01) //[10 0 0 0]

	//Arraylar ile oluşturulan dizler daha güvenlidir.
	//Daha hızlıdır

	//Arraylar kendi elemanlarını kopyalar. REferans değerlerini kopyalamaz. Js Referesan kopyalar...
	myArr3 := [3]int{1, 2, 3}
	fmt.Println(myArr3) //[1 2 3]
	myArr4 := myArr3
	fmt.Println(myArr4) //[1 2 3]

	myArr4[0] = 100
	fmt.Println(myArr4) //[100 2 3]

	fmt.Println(myArr3) //[1 2 3]

	fmt.Println("-------------------------------------")

	//Slicelar da refeerans değerini değiştirir.
	//Arraylar pass by value olarak çalışır
	//Slicelar ise Pass By Reference oalrak çalışır.

	mySlc02 := []int{1, 2, 3}
	fmt.Println(mySlc02) //[1 2 3]

	mySlc03 := mySlc02
	fmt.Println(mySlc03) //[1 2 3]

	mySlc03[0] = 32
	fmt.Println(mySlc03) //[32 2 3]

	fmt.Println(mySlc02) //[32 2 3]
}
