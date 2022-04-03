package main

import "fmt"

func main() {

	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(underArray)

	mySlc := underArray[2:5] //[2 3 4] --> 2 dahil 5 dağil değil
	fmt.Println(mySlc)

	mySlc01 := underArray[:]
	fmt.Println(mySlc01) //[0 1 2 3 4 5 6 7 8 9]

	mySlc02 := []int{3, 4, 5}
	mySlc02 = append(mySlc02, 6, 7, 9) //[3 4 5 6 7 9]     append() metodu ile slice a eleamn ekleyebiliyoruz.

	fmt.Println(mySlc02)

	mySlc03 := []int{5, 6, 7}
	mySlc04 := []int{8, 9}
	//	mySlc03 = append(mySlc03,mySlc04) //Burada bizlere hata verir. Bu hatayı aşabilmek için java daki varragslara benzer uygulamayı
	// kullanmamız gerekir

	mySlc03 = append(mySlc03, mySlc04...)
	fmt.Println(mySlc03)

	mySlcStr := []string{"a", "b", "c"}
	fmt.Println(mySlcStr) //[a b c]

	//Eleman nasıl silinir?

	mySlc05 := []int{1, 2, 3}
	mySlc05 = mySlc05[2:] //Baştaki ilk elemanı siler
	fmt.Println(mySlc05)  //[3]

}
