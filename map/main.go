package main

import "fmt"

func main(){
	var simpleMap = map[string]int{"hablu":10, "bablu": 20} // key will be Stirng and value Int
	fmt.Println(simpleMap)
	// Print map with key and value
	for key, value := range simpleMap{
		fmt.Printf("%s is %d \n", key, value)
	}

	// Print map value without key
	for _, value := range simpleMap{
		fmt.Printf("value: %d \n",value)
	}

	// Print map key without value
	for key := range simpleMap{
		fmt.Printf("key: %s \n",key)
	}

	result := map[string]string{
		"jhontu":"A+",
		"bantu": "A-",
		"jantu": "A",
	}
	
	fmt.Println("Result: ", result)
	
	// Add data in map
	result["pantu"]="B+"
	fmt.Println("new value added: ", result)

	// Delete item from map
	delete(result,"bantu")
	fmt.Println("After delete an element: ", result)

}