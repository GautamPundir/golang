package main
import "fmt"
func main (){
	a :=34
	b :=54
	result :=a+b
	fmt.Println("result :- ",result)
	x:=NumCPU()
	fmt.Println(x);
}
func NumCPU() int
