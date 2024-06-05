// recursão é a capcidade de uma fução chamar ela mesma
package main

import "fmt"

func fatorial(x int/*"int" mostar que x deve ser inteiro*/)int/*"int mostra que a função deve retornar um inteiro"*/{
	if x == 0 {
		return 1
	}

	return x * fatorial(x - 1)
}

func main() {
	fmt.Println(fatorial(9))
	fmt.Println(fatorial(2))
}