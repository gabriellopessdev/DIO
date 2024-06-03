// Estruturas são coleções de "campos" usadas para agrupar dados e formar registros
package main

import "fmt"

type pessoa struct{
	nome string
	idade int
	sexo string
}

func main() {

	fmt.Println(pessoa{"Gabriel", 20, "M"})
	fmt.Println(pessoa{"Ana", 28, "F"})
	
}