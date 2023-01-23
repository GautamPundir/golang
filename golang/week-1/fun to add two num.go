package main
import "fmt"
func  main(){
	a := 23
	b := 34
	c :=  sum(a,b)
	fmt.Printf("(using printf  )Sum of %d and %d is : %d\n",a,b,c)
	fmt.Println("(using println)Sum of a and b is :",c)
}
func sum(x int,y int) int{
	return x+y
}
