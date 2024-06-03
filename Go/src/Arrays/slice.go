package main

import "fmt"

func main() {

	array := [7]float64 {1,2,3,4,5,}
	slice1 := array[2:5]//selecionando uma fatia do array
	slice2 := append(slice1, 9,8,8,8,8,8,7)//"append" para adicionar elementos em uma fatia
	//"copy" para copiar uma fatia para outra
	fmt.Println(slice1, slice2)

}