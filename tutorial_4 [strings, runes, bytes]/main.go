package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("%v, %T", indexed, indexed)
	for i, v := range myString{
		fmt.Println(i, v)
	}
	fmt.Printf("\n the length of string is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\n myRune = %v", myRune)

	var strSlice = []string{"s","u","b","s","c","r","i","b","e"}
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}