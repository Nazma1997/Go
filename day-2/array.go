package main
import "fmt"

func main(){

	// Array declaration  type 1
	arrayType1 := [...]int{1,2,3,4}
	fmt.Println(arrayType1)

	// Array declaration  type 2
	arrayType2 := [...]int{1,2,3,4}
	fmt.Println( arrayType2)

	// Array declaration  type 3
	arrayType3 := [4]int{1,2,3,4}
	fmt.Println( arrayType3)

	// access array element
	accessArrayElement := [...]int{1,2,3,4,5}
	fmt.Println(accessArrayElement[1])


	// change array element

	changeArrayElement := [...]int{1,2,3,4,5}
	changeArrayElement[1] = 20
    fmt.Println(changeArrayElement)


	//Specific array element
	specificArrayElement := [...]int{2: 10, 3: 20}
	fmt.Println(specificArrayElement)

}

