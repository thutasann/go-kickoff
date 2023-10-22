package main

import "fmt"

type Car struct {
	model string
}

func main(){
	
	c := make(chan Car, 3)

	go func(){
		c <- Car{model: "model 1"}
		c <- Car{model: "model 2"}
		c <- Car{model: "model 3"}
		c <- Car{model: "model 4"}
		c <- Car{model: "model 5"}
		close(c)
	}()


	for i := range c {
		fmt.Println(i.model)
	}

}