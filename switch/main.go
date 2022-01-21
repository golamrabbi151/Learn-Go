package main
import (
    "fmt"
    "time"
)

func main(){
	i := 2
// Example one
	switch i {
	case 2:
		fmt.Println("Hello. The received value is", i)
	case 3:

	}
// Example two
	switch time.Now().Weekday(){
	case time.Friday, time.Saturday :
		fmt.Println("Hi, This is weekend")
	default:
		fmt.Println("It's a regular day")
	}
	
	// function

	whatAmI := func (i interface{})  {
		switch t:=i.(type){
		case int:
			fmt.Println("I am Integer")
		case bool:
			fmt.Println("Hello Boolean")
		default:
			fmt.Printf("Don't not type %T \n", t)

		}
	}

	whatAmI(true)
	whatAmI(20)
	whatAmI("hello")
	
}