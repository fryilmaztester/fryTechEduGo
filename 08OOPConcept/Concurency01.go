package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//WaitGroup üç önemli metot vardır. --> Add(),Done(),Wait()

func main() {

	//Go da eş zamanlı olarak yapılan işlere corruncey denir.
	//

	wg.Add(1) //--> ana go routine senin bekleyecğein n  tane go routine vardır.
	go printX()
	wg.Wait()
	fmt.Println()
	printY()
	time.Sleep(time.Second)
}

func printX() {

	for i := 0; i < 20; i++ {
		fmt.Print("X ")
	}
	wg.Done()
}

func printY() {

	for i := 0; i < 20; i++ {
		fmt.Print("Y ")
	}
}
