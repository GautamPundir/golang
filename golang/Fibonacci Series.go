package main
import "fmt"
func main()  {
	var a int  =0
	var b int  =1
	var c int  =0
	var n int
	fmt.Println("ENTER THE NO. OF TERMS OF FB SERIES")
	fmt.Scanln(&n)
	for i:=0; i<=n-1 ;i++{
		if (i==0){
			fmt.Println(a)
		}
		if (i==1){
			fmt.Println(b)	
		}else{
			c=a+b
			a=b
			b=c
			fmt.Println(c)
		}
	}
}
