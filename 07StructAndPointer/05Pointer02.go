/*package main

import "fmt"

func main() {

	x := 5  GO PASS BY VALUE Çalışır.
	fmt.Println(x) //5
	double(x)      // num 10
	fmt.Println(x) //5
}

func double(num int) {

	num *= 2
	fmt.Println("num: ", num)
}

/*package main

import "fmt"

func main() {

	//FAKAT SLC ler PASS BY REFERENCE ÇALIŞIR

	mySlc := []int{1, 10, 100}
	fmt.Println(mySlc) //[1 10 100]

	double(mySlc) //[2 20 200]

	fmt.Println(mySlc) //[2 20 200
}

func double(slc []int) {

	for i := 0; i < len(slc); i++ {
		slc[i] *= 2
	}

	fmt.Println(slc)
}*/

package main

import "fmt"

func main() {

	x := 5
	fmt.Println(x) //5
	//double(&x)     //0xc0000a6058
	double(&x)
	fmt.Println(x)

}

func double(num *int) {

	fmt.Println(num)
	*num *= 2
	fmt.Println(*num)
}
