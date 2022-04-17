package main

import "fmt"

type rectangle struct {
	a, b float64
}

func (r rectangle) area() float64 {

	return r.a * r.b
}

func (r rectangle) circulum() float64 {

	return 2 * (r.a + r.b)
}

type shape interface {
	area() float64
	circulum() float64
}

func interfaceFunc(i shape) {

	fmt.Println(i)
	fmt.Println("i ye ait olan area metodu: ", i.area())
	fmt.Println("i ye ait olan cevre metodu: ", i.circulum())

	fmt.Printf("%T %v\n", i, i)
}

func main() {

	//Bazı durumlarda bizim için önemli olan şey metotlardır.
	//Interface sadece o fonksiyonalrın imzasına sahiptir. Çalıştırma işlemi yapmaz.

	r1 := rectangle{
		3, 8,
	}
	/*fmt.Println(r1.area())
	fmt.Println(r1.circulum())*/

	interfaceFunc(r1)
}
