/*package main

import "fmt"

type myType int

func main() {

	//1) "int" veri tipi olacak şekilde kendi veri tipinizi oluştrunuz.
	var x myType
	x = 10
	fmt.Printf("%T %v", x, x) //main.myType 10




}*/

/*package main

import "fmt"

func main() {

	//2) Başlangıç değeri 10 olanbir x değişkeni oluşturnuz
	//Bulunduğu adres üzerinde ydeğişkeni tanımlayıpx değerini 20 yapınız.

	x := 10
	fmt.Println(x) //10
	y := (&x)
	*y = 20
	fmt.Println(x) //20

}*/

/*package main

import "fmt"

type myRectangle struct {
	short int
	long  int
}

func main() {

	//3) Rectangele bir struct type oluşturunuz
	//display, area,circumReference metodlarını oluştrunuz

	r1 := myRectangle{
		3, 5,
	}
	display(r1)
	fmt.Println("Area: ", area(r1))
	fmt.Println("Çevre: ", circulum(r1))
}

func display(r myRectangle) {

	fmt.Println("Kenar Uzunlukları: ", r.short, " ve ", r.long, " olan dikdörtgendir")
}

func area(r myRectangle) int {

	return r.short * r.long
}

func circulum(r myRectangle) int {

	return 2 * (r.short + r.long)
}*/

package main

import "fmt"

type myFamily struct {
	name      string
	age       int
	isMarried bool
}

func showFamily() []myFamily {

	f1 := myFamily{
		"Fatih",
		21,
		false,
	}

	f2 := myFamily{
		"Eda",
		32,
		true,
	}

	f3 := myFamily{
		"Yaşar",
		3,
		false,
	}

	return []myFamily{f1, f2, f3}
}

func main() {

	//4)4 adte user ıstruct yapısıylafarklı şekilde tanumlanyınız.name, age int
	//for döngüsü ile kullanıları gösteriniz.

	families := showFamily()
	for i := 0; i < len(families); i++ {

		fmt.Printf("%T %v\n", families[i], families[i])
	}
}
