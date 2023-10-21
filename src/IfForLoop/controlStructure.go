package main

import (
	"fmt"
)

func main(){
	f := true
	flag := &f

	// If / Else if
	if flag == nil || 2 == 3 {
		fmt.Println("Value is nil")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}


	// For loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for { Infinite Loop
	// 	fmt.Println("Love you babyyyyyy")
	// }

	arr := []string{"thuta", "hsu", "pyae"}
	for i, value := range arr {
		fmt.Println("index => ", i);
		fmt.Println("value => ", value)
	}

	myMap := make(map[string]interface{})
	myMap["name"] = "thuta"
	myMap["age"] = 20;

	for k, v := range myMap {
		fmt.Printf("Key: %s and Value: %v", k ,v);
	}	

}