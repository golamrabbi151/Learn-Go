package main

import "fmt"

func main ()  {
	var a int = 10
	b:= "hello world"
	var c string

	if a ==20 {
		fmt.Println("a: ", a)
	}else {
		fmt.Println("a not equal 20")
	}


	if (b == "hello world") {
		fmt.Println(b)
	}

	if c == "" {
		fmt.Println("C is a empty string")
	}

}