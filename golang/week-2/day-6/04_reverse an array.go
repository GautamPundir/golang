package main
import "fmt"
func main(){
	arr:= [5]int{9,6,5,7,4}
	var a int 
	var x =len(arr)-1
	fmt.Println("Array before reverse :",arr)
	for  i :=0;i<x/2;i++{
		a=arr[i]
		arr[i]=arr[x-i]
		arr[x-i]=a
	}
	fmt.Println("Array  after reverse :",arr)

}
