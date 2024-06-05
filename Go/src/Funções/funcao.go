package main

import "fmt"

func media(lista []float64)float64  {

	total := 0.0
	for _, valor := range lista {//valor pega(renge) slice lista
		total += valor
	}
	return total/float64(len(lista))
}

func main() { //Calcula media de uma sala

	lista := []float64{98, 93, 88, 55, 44}

	fmt.Println(media(lista))
}