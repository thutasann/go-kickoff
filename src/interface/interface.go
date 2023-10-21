package main

import "fmt"

// Car Interface
type ICar interface {
	Drive()
	Stop()
}

// ------- Labmbo struct
type Lambo struct {
	lamboModel string
}

func (l *Lambo) Stop(){
	fmt.Println("Stopping Lambo...")
}

func NewModel(arg string) ICar {
	return &Lambo{arg}
}

// ------- Chevy struct
type Chevy struct {
	chevlModel string
}

// Lambo Drive func
func (l *Lambo) Drive() {
	fmt.Println(l.lamboModel, " : Lambo is moving")
}

// Chevy Drive func
func (c *Chevy) Drive() {
	fmt.Println(c.chevlModel, " : Chevy is moving")
}

// Main Func
func main() {
	l := Lambo{lamboModel: "Gallardo"}
	c := Chevy{chevlModel: "C369"}
	l.Drive()
	c.Drive()
}
