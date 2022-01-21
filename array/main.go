package main

import ("fmt")

func main(){
	// Declear and initilize
	a := [5]int{1,2,3,4,5}
	fmt.Println(a)

	// Declear
	var b [5]int
	fmt.Println("Initial value",b)
	b[0]=1
	b[1]=2
	b[2]=3
	b[3]=4
	b[4]=5
	// Use Printf() with new lint (\n)
	fmt.Printf("Insert value in array: %d \n",b)

	var c []int
	fmt.Println("Empty array: ", c)
	// Insert value in empty array
	c = []int {10,20,30,40,50}
	fmt.Println("C is not empty now: ", c)

	// Slice, Unser the hood slice is a array 
	d:= make([]int, 4)
	fmt.Println(d)
	d = append(d,20,21)
	fmt.Println("Appen value in array: ", d)

	// Multiple slice
	d[1] = d[len(d)-1]
	fmt.Println("D: ",d)
	fmt.Println("D: ", d[0:3])

	for key, value := range d {
		fmt.Printf("%d is %d \n", key, value)
	}
}