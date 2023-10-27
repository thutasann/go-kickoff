package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, quits chan struct{}){
	// switch case
	// select for channsl async
	for  {
		time.Sleep(time.Second)
		select {
		case <- c :
			fmt.Println("Received => ", c)
		case <- quits :
			fmt.Println("Quit => ", quits)
			os.Exit(1);
		}
	}
}

func main(){
	c := make(chan int)
	quits := make(chan struct{})
	go Select(c, quits)

	go func(){
		c <- 1;
		quits <- struct{}{}
	}()

	select {}
}