package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {

	//fmt.Println(evenNum(4))

	result, err := evenNum(10)
	if err != nil {
		fmt.Println("Hata: ", err)
	} else {
		fmt.Println("Girilen Sayı: ", result)
	}

	fmt.Println("--------------------------------")

	//fmt.Println(sRoot(-2))
	resultTwo, errTwo := sRoot(-4)

	if errTwo != nil {
		fmt.Println(errTwo)
	} else {
		fmt.Println(resultTwo)
	}

	fmt.Println("--------------------------------")

	file, err := os.Open("05FunctionsAndException/test.txt")
	if err != nil {
		fmt.Println("Hata: ", err)
	} else {
		fmt.Println("Dosyamız: ", file)
	}

}

func evenNum(num int) (int, error) {

	if num%2 != 0 {
		return 0, errors.New("Exceptionnn, You can not enter even number")
	}

	return num, nil

	//nil --> Başlangıç değeri olmayan ifadelere verilen değerdir.

}

func sRoot(num float64) (float64, error) {

	if num < 0 {
		return 0, errors.New("Exception, NEGATIVE Sayıalrın Karekökü alınamaz")
	}
	return math.Sqrt(num), nil

}
