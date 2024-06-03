package main

//Desafio para exibir "Pin" todos os numeros de 0 a 100 Multiplos de 3 "Pan" para os Mutiplos de 5 e "Pin Pan" para Multiplos de ambos


import "fmt"

func main() {
	for number := 0; number < 100; number++ {
		switch{
		case number %3 == 0 && number %5 == 0:
			fmt.Println("Pin Pan")
		
		case number %3 == 0:
			fmt.Println("Pin")
			
		case number %5 == 0:
			fmt.Println("Pan")
			
		default:
			fmt.Println(number)
		}

	}
}