package main

import "fmt"

type retangulo struct {//estrutura
	comprimento, altura int
}

func (calcula retangulo) area() int {//metodo area
	return calcula.comprimento * calcula.altura 
}

func (calcula retangulo) perimetro() int {//metodo perimetro
	return 2*calcula.comprimento + 2*calcula.altura

}

func main() {
	calcula := retangulo{comprimento: 10, altura: 20}
	fmt.Println("Area do Retangulo:", calcula.area())//chamando metodo area
	fmt.Println("Perimetro do Retangulo:", calcula.perimetro())//chamando metodo perimetro

}
