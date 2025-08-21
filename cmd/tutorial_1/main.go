package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum += 1
	fmt.Println(intNum) 

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString1 string = `Hello
	World`
	fmt.Println(myString1)

	var myString2 string = "Hello" + "-" + "World"
	fmt.Println(myString2)

	fmt.Println(utf8.RuneCountInString("empat"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBool bool = false
	fmt.Println(myBool)

	var intNum3 rune
	fmt.Println(intNum3)

	myVar := "text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const pi float32 = 3.1415
}