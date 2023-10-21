package main

import "fmt"

func main() {

	flag := true;

	// continue / break
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flag = false
			break
		} else if i == 1{
			continue;
		}
		fmt.Println(i)
	}
	fmt.Println(flag)

	// Switch
	day := "Fri"
	
	switch day {
	case "Fri":
		fmt.Println("TGIF")
	case "Mon":
		fmt.Println("Oh this is monday")
	case "Tue":
		fmt.Println("This is Tuesday")
	default :
		fmt.Println("This is default")
	}

	switch {
	case day == "Fri" :
		fmt.Println("This is Friday")
	default :
		fmt.Println("This is Default")
	}
}
