package main

import "fmt"

func main (){
	fmt.Println("Starting of pointer")
	// var ptr *int
	// fmt.Println("The pointer is : ",ptr)
	// fmt.Printf("The data type of ptr is %T",ptr)
	var value = 24
	ptr := &value
	fmt.Println("The value is ",*ptr)
	fmt.Println("The pointer address is ",ptr)
	*ptr= *ptr+2;
	fmt.Println("Updated value is ",*ptr)
	fmt.Println("Pointer address after updatation in value ",ptr)
}
