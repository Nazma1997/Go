package main

import "fmt"

func main(){
	mySlice := make([]int, 5, 10)

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}