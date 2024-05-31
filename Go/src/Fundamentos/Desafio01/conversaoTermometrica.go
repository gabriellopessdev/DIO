package main

import "fmt"

const ebulicao = 373

func main() {
	tempK := ebulicao
	tempC := tempK - 273

	fmt.Println("A temperatura de ebulição da água em Kelvin é de", tempK, "e em graus Celcius é de", tempC)
}