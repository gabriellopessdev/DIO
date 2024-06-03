package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perimetro() float64
}

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return math.Pow(q.lado, 2)}

func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (c circulo)  perimetro() float64{
	return 2 * math.Pi * c.raio
	
}

func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	q := quadrado{lado :5}
	c := circulo{raio :10}

	medir(q)
	medir(c)
}