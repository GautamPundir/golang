package main
import ("fmt"
		"time"
	)
func main(){
	present := time.Now()
	fmt.Println("current time is : ",present)
	fmt.Println("date",present.YearDay())
	fmt.Println("month",present.Month())
	fmt.Println("Year",present.Year())
	fmt.Println("Hour",present.Hour())
	fmt.Printf("%d-%d-%d \n",present.Day(),present.Month(),present.Year())
	fmt.Println(present.Format( "2-1-2006" ))
	fmt.Println(present.Format( "2006-1-2006" ))
	fmt.Println(present.Format( "02-01-2006 Monday" ))