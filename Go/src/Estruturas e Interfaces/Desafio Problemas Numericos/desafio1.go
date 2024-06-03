package main

//Desafio para exibir todos os numeros de 0 a 100 divisiveis por 3

func main() {
	
	for number := 0; number <= 100; number++ {
		if number %3 == 0 {
			println( number)
		}
		
	}
}