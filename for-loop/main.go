package main

import "fmt"

func main(){
	// Simple for loop
	fmt.Println("Simple for loop")
	for i := 0; i < 5; i++ {
		fmt.Println("number: ", i)
	}

	// Simple while loop
	 n := 1
fmt.Println("Simple while loop")
	for n<5{
		println("number: ", n)
		n++
	}

// Infinite loop

	for {
		n++
		if n==10 {
			fmt.Println("Infinite loop brake with value n:10")
			break
		}
	}

}