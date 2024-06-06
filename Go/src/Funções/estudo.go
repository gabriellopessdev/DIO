package main
import "fmt"
//Funções com paramentro opcional

func torrar(pao string,valor float64, nome ... string){
     cliente := "Não teve cliente indentificado"//variavel para trtar paramentro opcinal representado por "nome ..."
    switch {
 	case len(nome) > 0: // se nome tiver um parametro indentificado "cliente = nome[0]"
         	cliente = nome[0]
     		fmt.Println("Torrada feita com", pao)
     		fmt.Println("é um pedido de", cliente)
			fmt.Println("Valor total é", valor)
 	default:
 		fmt.Println("Torrada feita com", pao)
 		fmt.Println(cliente)
		fmt.Println("Valor total é", valor)

	}
}

// func torrar(pao string, valor float64, nome ... string){
//     cliente := "Não teve cliente indentificado"

//     if len(nome) > 0 {
//         cliente = nome[0]
//         fmt.Println("Torrada feita com", pao)
//         fmt.Println("é um pedido de", cliente)
// 		fmt.Println("Valor total é", valor)
// 		} else {
// 			fmt.Println("Torrada feita com", pao)
// 			fmt.Println(cliente)
// 			fmt.Println("Valor total é", valor)
//     }
// }

func main() {
    torrar("pão de forma", 12.99, "Gabriel")
    torrar("pão de sal", 10.99)

}
