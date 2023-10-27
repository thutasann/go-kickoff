package main

import (
	"fmt"
	"time"
)

// The pinger prints a ping and waits for a pong (indefinitely running in the background)
func pinger(pinger <- chan int, ponger chan <- int){
	for {
		<- pinger
		fmt.Println("pingggg");
		time.Sleep(time.Second)
		ponger <- 1
	}
}

// The ponger prints a pong and waits for a ping (indefinitely running in the background)
func ponger(pinger chan <- int, ponger <- chan int){
	for {
		<- ponger
		fmt.Println("pongggg");
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main(){
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	// The main goroutine starts the ping/pong by sending into the ping channel
	ping <- 1

	// for {
	// 	// Block the main thread until an interrupt
	// 	time.Sleep(time.Second)
	// }

	select {}
}