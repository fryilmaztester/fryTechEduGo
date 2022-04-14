package main

import "fmt"

func main() {

	//Pointer başka bir değişkenin adresini tutar

	name := "arin"
	fmt.Println(name)

	fmt.Println(&name) //0xc00010e110 ---->>>>name==>  değişkenin bulundunğu hafızadaki adresi
	//& --> adres operatoru. O değikenin adresini bizlere verir

	x := 22
	fmt.Println(x)
	fmt.Println(&x) //0xc0000160f0 --> x değişkenin bulunduğu hafızadki adresi

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", &x, &x) //*int, 0xc0000aa0a0 --> Pointerin veri tipi * lidir.

	/*
		var y int = &x
		fmt.Println(y) //Burada bizler hata alırız. NİYE??? x in veri tipi pointer yani * lidir. y ninki ise int olduğu için data type lar match olmaz.
	*/

	y := &x
	fmt.Println(y)
	fmt.Printf("%T, %v\n", y, y) //*int, 0xc0000160f0

	//* operatorun farklı bir işlevi daha vardır.
	fmt.Println(x)     //22
	fmt.Println(&x)    //22 nin adresi
	fmt.Println(*(&x)) //Dereferencing. * operatoru parantez içinde bulunan adresteki value return olur

	y1 := 10
	y2 := &y1

	fmt.Println(y1, y2)  //10 0xc000016120
	fmt.Println(y1, *y2) //10 10

	*y2 = 3
	fmt.Println(y1, *y2) //3 3

	y3 := 10
	y4 := y3
	fmt.Println(y3, y4) //10 10
	y4 = 8
	fmt.Println(y3, y4) //10 8
	y3 = 7
	fmt.Println(y3, y4) //7 8

	//z1 := [4]int{1, 10, 100, 1000} //array --> pass by value

	z1 := []int{1, 10, 100, 1000} // slice ---> pass by reference
	z2 := z1

	fmt.Println(z1, z2) //[1 10 100 1000] [1 10 100 1000]

	z2[0] = 3
	fmt.Println(z1, z2) //array pass by value oalrak çalıştığı için value değişmez. [1 10 100 1000] [3 10 100 1000]

	//slice ---> pass by reference çalıştığı için value değişir. Referance ı değiştirir....[3 10 100 1000] [3 10 100 1000]

}
