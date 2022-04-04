package main

import (
	"fmt"
)

func main() {

	myMap := map[string]int{
		"Ahmet": 40,
		"Fatih": 21,
		"Selin": 32,
		"Meral": 28,
	}

	fmt.Println(myMap) //map[Ahmet:40 Fatih:21 Meral:28 Selin:32]

	//Specifick Elemanı nasıl yazdırırız.
	fmt.Println(myMap["Ahmet"]) //Burada Key değerini vermemiz gerekiyor. Value verir isek kabul etmez. --> 40

	myMap01 := map[string]int{
		"Ahmet": 40,
		"Fatih": 21,
		"Selin": 32,
		"Meral": 28,
	}

	fmt.Println(myMap01)

	//Map in içinde bulunmayan Key değerinin valuesına ulaşmak istediğimiz zaman GO bize value değerinin zzero değerini döner
	fmt.Println(myMap01["Aslı"]) //0

	isMarried := map[string]bool{
		"Fatih":  false,
		"Ali":    true,
		"Hatice": false,
		"Emine":  true,
	}

	fmt.Println(isMarried) //map[Ali:true Emine:true Fatih:false Hatice:false]

	fmt.Println("---------------------")
	//MAp Uzunluğu Bulma
	fmt.Println(len(myMap01)) //4

	fmt.Println("-------------------------")
	//Bir verinin map te olup olmadığını nasıl bulabiliriz
	value, ok := isMarried["Fuat"] //false false
	fmt.Println(value, ok)         //ok --> Ya true ya da false dönecektir. Burada değer varsa true, yoksa false döner. Value değeri ise
	//Eğer eleman yok ise value değerinin zero değeri var ise o eleman bize döner.
	value, ok = isMarried["Ali"]
	fmt.Println(value, ok) //true true

	if ok == true {
		fmt.Println("Bu değer map te bulunmaktadır. ")
	}

	fmt.Println("----------------------------")

	//Nasıl silebiliriz
	delete(isMarried, "Fatih")
	fmt.Println(isMarried)

	//MAKE Methodu ile Map Creation

	myMap02 := make(map[string]int)

	fmt.Println(myMap02) //map[]

	//MAPLER DE Sliceler gibi PASS By REference olarka çalışır.

}
