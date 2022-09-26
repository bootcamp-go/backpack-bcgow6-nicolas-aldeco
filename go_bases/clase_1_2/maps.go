package main

import "fmt"

var miVariable interface{}

func main() {
	var myOtherArray = [...]int{1, -2, 3, 4, 34, 123}
	fmt.Println("myOtherArray", myOtherArray)
	slice := myOtherArray[1:3]
	fmt.Println("slice", slice)
	fmt.Println("len", len(slice))
	fmt.Println("cap", cap(&myOtherArray))

}



