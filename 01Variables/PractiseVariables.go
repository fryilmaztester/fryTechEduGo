package main

import "fmt"

func main() {

	//1) studentName --> John Doe,   grade --> 77,isPassed --> true
	//3 farklı yontemle ile olusturup,consola yazdıırnız

	//1). Method

	/*var studentName string = "John Doe"
	var grade int = 77
	var isPassed bool = true

	*/

	//2.Method

	/*var studentName = "John Doe"
	var grade = 77
	var isPassed = true*/

	//3. Method

	/*	studentName := "John Doe"
		grade := 77
		isPassed := true*/

	fmt.Println("--------------------------------------------------------------------------------------------------")

	//2) Yukarıda belirtilen değişkenleri tek satır icinde tanımlayınız

	//1.Method

	/*var (
		studentName = "John Doe"
		grade       = 77
		isPassed    = true
	)*/

	//2.Method
	/*var studentName, grade, isPassed = "John Doe", 77, true*/

	//3.Method
	studentName, grade, isPassed := "John Doe", 77, true

	fmt.Println("studentName: ", studentName)
	fmt.Println("grade: ", grade)
	fmt.Println("isPassed: ", isPassed)
}
