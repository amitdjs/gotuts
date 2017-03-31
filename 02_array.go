package main

import "fmt"

func main () {
	var arrs[5] int
	var i int

	for i = 0; i<5; i++ {
		arrs[i] = i*20
	}

	for i=0; i<5; i++ {
		fmt.Println(arrs[i])
	}
}
