package main

import "fmt"

func main() {

	//Sadece for loop bulunmaktadır....

	for i := 1; i < 10; i++ {
		fmt.Print(i, " ") //1 2 3 4 5 6 7 8 9
	}

	fmt.Println()
	fmt.Println("------------------------")

	k := 0
	for ; k < 10; k += 5 {
		fmt.Print(k, " ") //0 5
	}

	fmt.Println()
	fmt.Println("----------------------------")
	/*
		for {
			fmt.Println("FRY") // ınfinite loop denir...
		}*/

	fmt.Println()
	fmt.Println("----------------------------")

	j := 10
	for j >= 0 {
		fmt.Print(j, " ") //10 9 8 7 6 5 4 3 2 1 0
		j--
	}

	fmt.Println()
	fmt.Println("----------------------------")

	for m := 0; m <= 10; m++ {
		if m%3 == 0 {
			continue
			//Eger şart sağlanırsa continue çalışır yani, çalıştığı yerde tekrardan loop un başına döner
		}
		fmt.Print(m, " ")
	}

	fmt.Println()
	fmt.Println("----------------------------")

	for d := 0; d <= 10; d++ {
		if d == 3 {
			break
			//Eger şart sağlanırsa break çalışır yani, çalıştığı yere kalan olan kısım yazdırılır
		}
		fmt.Print(d, " ")

	}
}
