package main

import "fmt"

type Car struct {
	name    string
	age     int
	modelNo int
}

// Print Func
func (c Car) PrintCar(){
	fmt.Println(c)
}

// Drive Func
func (c Car) Drive(){
	fmt.Println("Car is Driving...")
}

// Get Name Func
func (c Car) getName(){
	fmt.Println("here is the car name", c.name)
}

func main() {
	car := Car{
		name:    "toyota",
		age:     2002,
		modelNo: 2003,
	}
	car.PrintCar();
	car.Drive();
	car.getName();
}
