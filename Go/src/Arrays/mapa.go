package main

import "fmt"

func main() {

	x := make(map[string]int) //"x" é um mapa de string para int
	x["chave"] = 10
	x["teste"] = 5

	fmt.Println(x["chave"], x["teste"])
}
