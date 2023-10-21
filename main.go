package main

func swap(m1, m2 *int){
	var temp int
	temp = *m2;
	*m2 = *m1;
	*m1 = temp
}

// func main(){
// 	var m1 string
// 	m1 = "thuta sann"
// 	fmt.Println(strings.ReplaceAll(m1, "t", "NO"))
// 	fmt.Println(strings.Split(m1, " "))
// }

// func main()  {
// 	m1 := 1;
// 	m2 := 2;
// 	fmt.Println(int(m1) + m2);
// }