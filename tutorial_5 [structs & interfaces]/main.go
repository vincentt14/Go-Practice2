package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type owner struct{
	name string
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh*e.mpkwh
}

type engine interface{
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8){
	if miles <= e.milesLeft(){
		fmt.Println("You can make it")
	}else {
		fmt.Println("Need more fuel up frist")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15}
	myEngine.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())

	var myEngine2 electricEngine = electricEngine{100, 35}

	canMakeIt(myEngine, 50)
	canMakeIt(myEngine2, 100)
}