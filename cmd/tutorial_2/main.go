package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello world"
	printMe(printValue)

	var nuemrator = 11
	var denominator = 2
	var result, remainder, err = intDivision(nuemrator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	}else if remainder == 0 {
		fmt.Printf("The result of the integer is %v", result)
	}else {
		fmt.Printf("The result of the integer is %v with remainder %v", result, remainder)
	}

	switch remainder{
	case 0:
		fmt.Println("The Division was exact")
	case 1,2:
		fmt.Println("The Division was close")
	default:
		fmt.Println("The Division was not close")
	}
}

func printMe(printValue string) {
	fmt.Println("hello" + printValue)
}

func intDivision(nuemrator int, denominator int) (int, int, error) {
	var err error
	if denominator==0{
		err = errors.New("Cannot Devide by Zero")
		return 0,0, err
	}
	var result int = nuemrator/denominator
	var remainder int = nuemrator%denominator
	return result, remainder, err
}