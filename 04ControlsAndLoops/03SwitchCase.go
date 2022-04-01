package main

import "fmt"

func main() {

	switch grade := 3; grade {

	case 1:
		fmt.Println("Çok Kötü")
		break
	case 2:
		fmt.Println("Kötü")
		break
	case 3:
		fmt.Println("Orta")
		break
	case 4:
		fmt.Println("İyi")
		break
	case 5:
		fmt.Println("Çok İyi")
		break
	default:
		fmt.Println("Wrong Value.....")
		//Beklenen den farklı bir değer girildiği zaman bu değeri bize defeault olarak verir...
	}

	switch note := 6; note {

	case 1:
		fmt.Println("Çok Kötü")
		break
	case 2:
		fmt.Println("Kötü")
		break
	case 3:
		fmt.Println("Orta")
		break
	case 4:
		fmt.Println("İyi")
		break
	case 5:
		fmt.Println("Çok İyi")
		break
	default:
		fmt.Println("Wrong Value.....") //... Beklenenden farklı bir değer girildiği için Wrong Value çıktısını görmeyi bekleriz...
		fmt.Println("note: ", note)
	}

	//fallthrough --> Eğer fallthroug var ise diğer case leri de kontrol eder.

}
