package main
import "fmt"
func main(){
	arr := [4]int{2,4,5,3}
	var arr2 [4] int
	var ptr *int
	ptr=&arr[0]
	fmt.Println("The value is ",*ptr)
	fmt.Println("The pointer address is ",ptr)
	for i:=0;i<len(arr);i++{
		arr2[i] =arr[i]
	}
	fmt.Println("Array 2 is ",arr2)

}
