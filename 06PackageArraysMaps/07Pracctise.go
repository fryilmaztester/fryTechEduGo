package main

import "fmt"

func main() {

	/*
		//1)5 str elemanından oluşan bir array ve5 int elemanından oluşan bir slice oluştuurp index numaraları ile birlikte gösteriniz

		cities := [5]string{"İstanbul", "Tahran", "Roma", "Belgrad", "Paris"}

		for index, value := range cities {
			fmt.Println(index, value)
		}

		cities01 := []string{"İstanbul", "Tahran", "Roma", "Belgrad", "Paris"}

		for index, value := range cities01 {
			fmt.Println(index, value)
		}

		cities02 := [...]string{"İstanbul", "Tahran", "Roma", "Belgrad", "Paris"}

		for index, value := range cities02 {
			fmt.Println(index, value)
		}

		//2) [0,1,2,3,4,5,6,7,8] slice ından [0,1,2,3] , [4,5,6], [6,7,8], [2,3,6,7] slicelarını oluştuunuzmy

		mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

		mySlc01 := mySlc[:4]
		fmt.Println(mySlc01)

		mySlc02 := mySlc[4:7]
		fmt.Println(mySlc02)

		mySlc03 := mySlc[6:]
		fmt.Println(mySlc03)

		var mySlc04 []int
		mySlc04 = append(mySlc[2:4], mySlc[6:8]...)

		fmt.Println(mySlc04)


	*/

	//3) slicelar için copy veassign metoduun farklılılarını açıklayınız.

	/*mySlc05 := []int{1, 2, 3}
	mySlc06 := make([]int, 2)

	fmt.Println(mySlc05)
	fmt.Println(mySlc06)

	copy(mySlc06, mySlc05) //1. slice, 2. slice kopyalama işlemi yapar.

	fmt.Println(mySlc05)
	fmt.Println(mySlc06)
	mySlc05[0] = 100

	fmt.Println(mySlc05)
	fmt.Println(mySlc06)*/

	/*mySlc05 := []int{1, 2, 3}
	mySlc06 := make([]int, 2)

	fmt.Println(mySlc05)
	fmt.Println(mySlc06)

	mySlc06 = mySlc05

	fmt.Println(mySlc05)
	fmt.Println(mySlc06)
	mySlc05[0] = 100

	fmt.Println(mySlc05)
	fmt.Println(mySlc06)*/

	//4) map gösterimi ile yazarlar ve kitapları for range ile gösteriniz

	myMap := map[string][]string{
		"Yaşar Kemal":     []string{"İnce Memed", "Yer Demir Gök Bakır"},
		"Sabahattin Ali":  []string{"Kuyucaklı Yusuf", "Kürk Mantolu Madonna", "Değirmen"},
		"Haruki Murakami": []string{"1Q84", "Dans Dans Dans", "Kumandanı Öldürmek"},
	}

	fmt.Println(myMap)
	fmt.Println(myMap["Yaşar Kemal"])
	fmt.Println(myMap["Haruki Murakami"])

	for key, value := range myMap {

		fmt.Println("Yazarımız: ", key)
		fmt.Println("Bazı Kitapları Aşağıdadır: ")

		for i, v := range value {
			fmt.Println("\t", i+1, v)
		}
	}
}
