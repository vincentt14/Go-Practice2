package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costo.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrice(websites[i], chickenChannel)
		go checkTofuPrice(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrice(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second*1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice<=MAX_CHICKEN_PRICE{
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrice(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second*1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice<=MAX_CHICKEN_PRICE{
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chikenChannel chan string, tofuChannel chan string){
	select{	
	case website := <-chikenChannel:
		fmt.Printf("\nEmail sent: Found a deal on chicken at %v", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEmail sent: Found a deal on tofu at %v", website)
	}
}