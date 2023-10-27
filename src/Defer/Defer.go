package main

import "fmt"

func greet(){
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
}

func main()  {
	greet()
}