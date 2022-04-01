package main

import "fmt"

func main() {

	//if <boolean> { code statement }

	x := 8

	if x%2 == 0 {
		fmt.Println(x, " çift sayıdır") // x çift sayıdır
	}

	fmt.Println("---------------------------------------------")
	//if boolean expression ı true döenrse code bloğu çalışır

	y := 7

	if y%2 != 0 {
		fmt.Println(y, " tek sayıdır") // y tek sayıdır
	}
	fmt.Println("---------------------------------------------")

	w := 54
	if w%2 != 0 {
		fmt.Println(w, " tek sayıdır")
	} else {
		fmt.Println(w, " çift sayıdır")
	}

	fmt.Println("---------------------------------------------")

	if !false {
		fmt.Println("Message Sent") //false un değili true olduğu için kod bloğu çalışır
	}

	fmt.Println("---------------------------------------------")

	num1 := 27

	if num1 == 0 {
		fmt.Println(num1, " 0 a eşittir")
	} else if num1%2 == 0 {
		fmt.Println(num1, " çift sayıdır")
	} else {
		fmt.Println(num1, " tek sayıdır")
	}

	fmt.Println("---------------------------------------------")

	num2 := -4

	if num2 < 0 {
		fmt.Println(num2, " Negative")
	} else if num2 > 0 {
		fmt.Println(num2, " Positive")
	} else {
		fmt.Println(num2, " Nötr")
	}
}
