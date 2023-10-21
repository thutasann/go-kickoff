package main

import (
	"fmt"
	"time"
)

// Heavy
func heavy(){
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy.....")
	}
}

// Super Heavy
func superHeavy(){
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super Heavy.....")
	}
}

func main(){
	go heavy();
	go superHeavy();
	fmt.Println("Main")
	select {}
}