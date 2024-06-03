package main

import "fmt"

func main() {

	var nota [5]float64
	nota[0] = 5.3
	nota[1] = 8
	nota[2] = 4.2
	nota[3] = 2.1
	nota[4] = 7.8

	var total float64 = 0
	for i := 0; i < len(nota); i++ {
		total += nota[i]
	}
	fmt.Println(total / float64(len(nota)))
}