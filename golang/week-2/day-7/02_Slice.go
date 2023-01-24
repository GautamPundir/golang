package main

import (
	"fmt"
	
)
func main(){
	var sfruit[] int
	var s int
	fmt.Println("hi check")
	for i:=0;i<5;i++{
		fmt.Println("Enter fruits name")
		fmt.Scanf("%d",s)
		sfruit=append(sfruit,s)
		fmt.Println("hi check 2")
	}
	
	fmt.Println(sfruit)
}
