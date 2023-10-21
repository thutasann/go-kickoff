package main

import (
	"fmt"
	"strings"
)

func main(){
	var m1 string
	m1 = "thuta sann"
	fmt.Println(strings.ReplaceAll(m1, "n", "🚀"))
	fmt.Println(strings.Split(m1, " "))
	m1 = "Thuta Sann"
}
