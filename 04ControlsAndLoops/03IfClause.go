package main

import (
	"fmt"
	"runtime"
)

func main() {

	x := 10

	if x := -5; x < 0 { // x in scope u sadece if statement kapsamında düşünülmüştür. if in haricinde kullanıalmayacak şekildedir.
		fmt.Println(x, " Negative")
	} else if x == 0 {
		fmt.Println(x, " Notr")
	} else {
		fmt.Println(x, " Positive")
	}

	fmt.Println(x) //10 --->>>>>  Burdaki x ise 7. satırda bulunan x in değeridir

	fmt.Println("---------------------------------------------")

	//Nested If

	if y := 3; y < 0 {
		fmt.Println(x, " Negative")
	} else {
		if y%2 == 0 {
			fmt.Println(y, " Even")
		} else if y == 3 {
			fmt.Println("BREAKKKK")
			runtime.Breakpoint()
		} else {
			fmt.Println(y, " Odd")
		}
	}

}
