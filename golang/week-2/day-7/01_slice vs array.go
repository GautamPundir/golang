package main
import "fmt"
func main(){
	var slice [] int
	var arr [6]  int
	slice=append(slice,3)
	arr[0]=2
	arr[1]=4
	arr[2]=5
	fmt.Println(slice)
	fmt.Printf("%T\n",slice)
	fmt.Println(arr)
	fmt.Printf("%T\n",arr)
}
