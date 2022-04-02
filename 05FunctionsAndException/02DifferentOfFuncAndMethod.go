package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	//fmt.Println(result(43))
	merhaba("Fatih", 26)

	//önce merhaba() metodu calışır sonrasında allta gelen code lar çalışır.
	name := "Ramazan"
	age := 27
	fmt.Printf("Adım %s, yaşım %d\n", name, age)

	var x int = 10
	var moment time.Time = time.Now() // Now() --> Farklı bir paketten çağrılan func tır.

	fmt.Println(x)
	fmt.Println(moment)

	fmt.Print("Lutfen sınav sonucunu giriniz...: ")
	reader := bufio.NewReader(os.Stdin) // Kullanıcıdan veri almak için
	value, _ := reader.ReadString('\n') //_ black identifier
	fmt.Println(value)

	bolum, kalan := div(104, 5)
	fmt.Println(bolum, kalan)
}

/*func result(grade int) string {

	//return keyword unden sonra herhangib aşka bir şey yazdırılamaz....


	if grade > 50 {
		return "gectiniz"
		fmt.Println("GECTINIZZz")
	} else {
		return "kaldınız "
		fmt.Println("KALDINIZZZ")
	}
}*/

func merhaba(name string, age int) {
	fmt.Printf("Adım %s, yaşım %d\n", name, age)
}

func div(bolunen int, bolen int) (bolum int, kalan int) {

	bolum = bolunen / bolen
	kalan = bolunen % bolen

	return bolum, kalan
}
