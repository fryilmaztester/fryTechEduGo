/*package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A Circle")
	//wg01.Done()
}
func (c circle) area() float64 {
	return math.Pi * c.r * c.r

	//go routine in ilişkilendirildiği func ta return var ise go routine ile ilişkilendirilemez.
}

//var wg01 sync.WaitGroup

func main() {

	wg.Add(1)

	c1 := circle{5}
	area1 := c1.area()
	fmt.Printf("%.2f\n", area1)
	//go c1.display()

	c1.display()
	//wg01.Wait()
}
*/

/*package main

import "fmt"

func merhaba() string {

	return "merhaba"
}

//Channel lar go routine lerin bir biri ile iletişmlerini sağlar
//Aynı zaman da bir go routine tarafından gönderilen bir değeri diğer go routine kullanılmadan önce gelmesini garanti eder.

func main() {

	fmt.Println(merhaba())
}*/

/*package main

import "fmt"

func merhaba(c1 chan string) {
	c1 <- "Merhaba"
}
func main() {

	/*var myChannel chan string
	myChannel = make(chan string)

	myChannel := make(chan string)
	go merhaba(myChannel)

	fmt.Println(<-myChannel)
}*/

package main

import "fmt"

func main() {

	chan01 := make(chan string)
	chan01 <- "fry"
	fmt.Println(<-chan01)
}
