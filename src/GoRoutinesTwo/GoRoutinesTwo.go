package main

import (
	"fmt"
	"sync"
)

// GoRoutines with WaitGroups
func main(){

	// wait groups
	wg := sync.WaitGroup{}

	// add, done
	wg.Add(2);
	go func(){
		fmt.Println("Hello")
		wg.Done();
	}()
	go func(){
		fmt.Println("World")
		wg.Done();
	}()
	
	// wait
	wg.Wait()
	fmt.Println("Main")

	wg.Wait()
	fmt.Println("Exit")
}