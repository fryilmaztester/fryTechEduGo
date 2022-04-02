package main

import (
	"fmt"
	"strings"
)

func main() {

	//Package: Benzer kod blokalrının içinde tuutlduğu yerdir
	//Bize hazır metot, func sunan depolardır.
	//Paketlerden bu tür metotları biz hazır alırız.

	//PAKETTEN NASIL ÇAĞIRACAĞIZ
	fmt.Println("Hello")
	//Burada yazan Println() metodu kullanabilmek için Go, bu metotun buluğu metodu import eder.

	//Paket nasıl indirilir?
	//https://pkg.go.dev/ isteidiğimiz paketi bu sitede bulabiliriz.

	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))

	fmt.Println(strings.Count("animatrix", "a"))       //2
	fmt.Println(strings.Replace("fatia", "a", "e", 1)) //fetih

}
