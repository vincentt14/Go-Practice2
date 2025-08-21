package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[0] = 123
	fmt.Println(intArr)
	fmt.Println(intArr[1:3])	

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	intArr2 := [...]int32{1,2,3}
	fmt.Println(intArr2)

	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	
	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 10)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"adam":23, "sarah":45}
	fmt.Println(myMap2["adam"])
	var age, ok = myMap2["adam"]
	if ok != false{
		fmt.Printf("umur adam %v", age)
	}else {
		fmt.Println("ga ada adam")
	}
}