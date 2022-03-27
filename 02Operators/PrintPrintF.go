package main

import "fmt"

var x = 10

var Y = 100 // Buradaki Y değişkenine diğer paketlerden de iletişim vardır demektri bu.

func main() {

	fmt.Println("Merhaba")  //AltSatıra boşluk koyar
	fmt.Printf("\nHello\n") //Boşluk koymaz bundan sonraki printte bunun yanına yazdırılacaktır.

	//Vısıbılıty

	fmt.Println(x)

	//Naming

	//İsimlendirme sade olmalı fakat anlaşılır olmalıdır.
	//CamelCase kuralı burada da geçerlidir.

	var custName = "Fatih"

	fmt.Println(custName)
}
