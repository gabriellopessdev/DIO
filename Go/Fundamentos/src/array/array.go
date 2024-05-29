package main

import "fmt"

func main() {
	var x [10] [10] int
	x[4] [0]  = 80
	x[9] [0] = 150
	fmt.Println(x)
}