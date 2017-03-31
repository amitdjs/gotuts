package main

import "fmt"

func main() {
	var num int = 10
	var numptr *int = &num

	fmt.Println(num)
	fmt.Println(numptr)
	fmt.Println(&num)
	fmt.Println(*numptr)
}
