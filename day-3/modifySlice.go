package main

import "fmt"

func main(){

	// access element of a slice
	mySlice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice1[2])

	// change element of a slice
	mySlice1[1] = 10
	fmt.Println(mySlice1)

	// Append element of a slice

	mySlice1 = append(mySlice1, 20, 30)
	fmt.Println(mySlice1)


	//Append one slice to  another slice

	mySlice2 := []int{100,200}

	appendTwoSlice := append(mySlice1, mySlice2...)
	fmt.Println(appendTwoSlice)
}