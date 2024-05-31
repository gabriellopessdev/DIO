package main

import "fmt"

func main() {
	//formas de declaração de variaveis
	var name string = "Gabriel"//declatando tipo
	var age = 20//tipo não declardoo
	versao :=  1.1//operedor curto(Gopher) "nome da variavel" + ":="
	fmt.Println("Meu nome é", name,"eu tenho", age, "anos!")
	fmt.Println("Esta é a versão", versao, "do meu codigo!")
	fmt.Printf("Meu nome é %s, eu tenho %d anos!\n", name, age)//Utilizando printform
	fmt.Printf("Meu nome é %s (%T), eu tenho %d (%T) anos!\n", name, name, age, age)//Utilizando printformat para mostrar o tipo da variavel
}