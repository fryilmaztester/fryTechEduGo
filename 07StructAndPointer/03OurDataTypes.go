package main

import (
	"fmt"
	"strings"
)

type mile float64 //--->> mile float64 veri tipi üzerine kurulmuş kendi Defined Data Type ımızdır.
type kilometer float64
type myString string

func main() {

	var m1 mile
	m1 = 3.2
	fmt.Println(m1)
	fmt.Printf("%T %v\n", m1, m1) //main.mile 3.2

	fmt.Println("----------------------------")

	f1 := 4.1
	fmt.Printf("%T, %v\n", f1, f1) //float64, 4.1

	fmt.Println("----------------------------")

	//fmt.Println(m1 + f1) //m1 in data type ı mile dır. Fakat f1 in veri tipi float64 olduğu için işlem yapamayız.

	fmt.Printf("%T, %v\n", (m1 + mile(f1)), (m1 + mile(f1))) //main.mile, 7.3 --> Type Conversation Yapılarak f1 mile a donuşturulmuştur.

	//	k1 := kilometer(7.8)

	myString := "Ramazan"
	fmt.Println(strings.ToUpper(myString)) //RAMAZAN

	m3 := mile(10)
	k2 := toKilometer(m3)
	fmt.Println(k2)

	k3 := kilometer(4.3)
	m4 := toMile(k3)
	fmt.Println(m4)
}

func toKilometer(m mile) kilometer {

	return kilometer(m * 1.6)
}

func toMile(k kilometer) mile {

	return mile(k * 0.62)
}
