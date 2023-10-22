package main

import "fmt"

// <name> chan <datatype>

func main(){
	c := make(chan int)

	// --- send in a goroutine
	go func(){
		c <- 1
	}()

	// --- sniff
	val := <- c
	fmt.Println("val 1 => ", val);

	// --- send in a goroutine
	go func(){
		c <- 2
	}()

	// --- sniff again
	val = <- c
	fmt.Println("val 2 => ", val);
	
	fmt.Println("c => ", c)
}