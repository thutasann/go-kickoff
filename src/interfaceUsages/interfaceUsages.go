package main

import (
	"fmt"
)

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	Anything("thuta sann")
	Anything(struct{ name string }{
		name: "thuta",
	})

	mymap := make(map[string]interface {})

	mymap["name"] = "thuta sann"
	mymap["age"] = 12

	fmt.Println(mymap)

}
